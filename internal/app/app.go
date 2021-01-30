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
)

type options struct {
	ConfigFile string
	WWWDir     string
	Version    string
}

// Option
type Option func(*options)

// SetConfigFile
func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

// SetWWWDir
func SetWWWDir(s string) Option {
	return func(o *options) {
		o.WWWDir = s
	}
}

// SetVersion
func SetVersion(s string) Option {
	return func(o *options) {
		o.Version = s
	}
}

// Init
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

	logger.Printf(ctx, "start service，run mode：%s，version：%s，pid：%d", config.C.RunMode, o.Version, os.Getpid())

	// Initialize unique id
	iutil.InitID()

	// Initialize short id
	iutil.InitShortID()

	// Initiate monitor
	monitorCleanFunc := InitMonitor(ctx)

	// Initiate logger
	loggerCleanFunc, err := InitLogger()
	if err != nil {
		return nil, err
	}

	// Initiate injector
	injector, injectorCleanFunc, err := injector.BuildInjector()
	if err != nil {
		return nil, err
	}

	// Initiate gRPC & gateway server
	serverCleanFunc := InitServer(injector)

	return func() {
		serverCleanFunc()
		injectorCleanFunc()
		loggerCleanFunc()
		monitorCleanFunc()
	}, nil
}

// InitMonitor
func InitMonitor(ctx context.Context) func() {
	if c := config.C.Monitor; c.Enable {
		// ShutdownCleanup set false to prevent automatically closes on os.Interrupt
		// and close agent manually before service shutting down
		err := agent.Listen(agent.Options{Addr: c.Addr, ConfigDir: c.ConfigDir, ShutdownCleanup: false})
		if err != nil {
			logger.Errorf(ctx, "agent monitor error: %s", err.Error())
		}
		return func() {
			agent.Close()
		}
	}
	return func() {}
}

func InitServer(injector *injector.Injector) func() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	return injector.Server.Setup(ctx)
}

// Run
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
		logger.Printf(ctx, "signal received: [%s]", sig.String())
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
	logger.Printf(ctx, "quit service")
	time.Sleep(time.Second)
	os.Exit(int(atomic.LoadInt32(&state)))
	return nil
}
