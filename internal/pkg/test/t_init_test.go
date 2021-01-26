package test

import (
	"context"

	"github.com/gomaglev/microshop/internal/pkg/config"
	"github.com/gomaglev/microshop/pkg/logger"
)

const (
	configFile = "../../../configs/config.toml"
)

func InitConfig() context.Context {
	// 初始化配置文件
	config.MustLoad(configFile)

	config.C.RunMode = "test"
	config.C.Log.Level = 2
	config.C.Gorm.Debug = false
	config.C.Gorm.DBType = "sqlite3"
	ctx := logger.NewTraceIDContext(context.Background(), "test")
	return ctx
}
