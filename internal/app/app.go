package app

import (
	"context"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/gomaglev/microshop/internal/app/injector"
	"github.com/gomaglev/microshop/internal/pkg/config"
	iutil "github.com/gomaglev/microshop/internal/pkg/util"
	"github.com/gomaglev/microshop/pkg/logger"

	"github.com/google/gops/agent"
	"google.golang.org/grpc"
)

type options struct {
	ConfigFile string
	WWWDir     string
	Version    string
}

// Option 定义配置项
type Option func(*options)

// SetConfigFile 设定配置文件
func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

// SetWWWDir 设定静态站点目录
func SetWWWDir(s string) Option {
	return func(o *options) {
		o.WWWDir = s
	}
}

// SetVersion 设定版本号
func SetVersion(s string) Option {
	return func(o *options) {
		o.Version = s
	}
}

// Init 应用初始化
func Init(ctx context.Context, opts ...Option) (func(), error) {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	config.MustLoad(o.ConfigFile)
	if v := o.WWWDir; v != "" {
		config.C.WWW = v
	}
	config.PrintWithJSON()

	logger.Printf(ctx, "服务启动，运行模式：%s，版本号：%s，进程号：%d", config.C.RunMode, o.Version, os.Getpid())

	// Initialize unique id
	iutil.InitID()

	// Initialize short id
	iutil.InitShortID()

	// 初始化服务运行监控
	monitorCleanFunc := InitMonitor(ctx)

	// 初始化日志模块
	loggerCleanFunc, err := InitLogger()
	if err != nil {
		return nil, err
	}

	// 初始化依赖注入器
	injector, injectorCleanFunc, err := injector.BuildInjector()
	if err != nil {
		return nil, err
	}

	// Grpc server initialization
	server, serverCleanFunc := InitServer(injector)

	// Grpc gateway initialization
	gatewayCleanFunc := InitGateway(injector, server)

	return func() {
		gatewayCleanFunc()
		serverCleanFunc()
		injectorCleanFunc()
		loggerCleanFunc()
		monitorCleanFunc()
	}, nil
}

// InitMonitor 初始化服务监控
func InitMonitor(ctx context.Context) func() {
	if c := config.C.Monitor; c.Enable {
		// ShutdownCleanup set false to prevent automatically closes on os.Interrupt
		// and close agent manually before service shutting down
		err := agent.Listen(agent.Options{Addr: c.Addr, ConfigDir: c.ConfigDir, ShutdownCleanup: false})
		if err != nil {
			logger.Errorf(ctx, "Agent monitor error: %s", err.Error())
		}
		return func() {
			agent.Close()
		}
	}
	return func() {}
}

func InitServer(injector *injector.Injector) (*grpc.Server, func()) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	return injector.Server.Setup(ctx)
}

func InitGateway(injector *injector.Injector, server *grpc.Server) func() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if !config.C.Gateway.Enable {
		return func() {
			logger.Infof(ctx, "grpc gateway is not enabled")
		}
	}
	return injector.Gateway.Setup(ctx, server)
}

// Run 运行服务
func Run(ctx context.Context, opts ...Option) error {
	var state int32 = 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc, err := Init(ctx, opts...)
	if err != nil {
		return err
	}

EXIT:
	for {
		sig := <-sc
		logger.Printf(ctx, "接收到信号[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			atomic.CompareAndSwapInt32(&state, 1, 0)
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	cleanFunc()
	logger.Printf(ctx, "服务退出")
	time.Sleep(time.Second)
	os.Exit(int(atomic.LoadInt32(&state)))
	return nil
}
