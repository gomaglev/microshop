package gorm

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gomaglev/microshop/internal/app/model/gorm/entity"
	"github.com/gomaglev/microshop/internal/pkg/config"
	"github.com/gomaglev/microshop/pkg/logger"

	// Import the driver
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDB 创建DB实例
func NewDB() (*gorm.DB, func(), error) {
	var (
		ctx    = context.Background()
		cancel context.CancelFunc
	)

	cfg := config.C
	var dsn string
	var db *gorm.DB
	var err error

	if t := cfg.Gorm.Timeout; t > 0 {
		ctx, cancel = context.WithTimeout(ctx, t)
		defer cancel()
	}

	switch cfg.Gorm.DBType {
	case "mysql":
		dsn = cfg.MySQL.DSN()
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, nil, err
		}
	case "sqlite3":
		dsn = cfg.Sqlite3.DSN()
		err = os.MkdirAll(filepath.Dir(dsn), 0777)
		if err != nil {
			return nil, nil, err
		}
	case "postgres":
		dsn = cfg.Postgres.DSN()
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, nil, err
		}
	default:
		return nil, nil, errors.New("unknown db")
	}

	if cfg.Gorm.Debug {
		db = db.Debug()
	}

	sqldb, err := db.DB()
	cleanFunc := func() {
		err := sqldb.Close()
		if err != nil {
			logger.Errorf(ctx, "Gorm db close error: %s", err.Error())
		}
	}

	if err = sqldb.Ping(); err != nil {
		return nil, cleanFunc, err
	}

	sqldb.SetMaxIdleConns(cfg.Gorm.MaxIdleConns)
	sqldb.SetMaxOpenConns(cfg.Gorm.MaxOpenConns)
	sqldb.SetConnMaxLifetime(time.Duration(cfg.Gorm.MaxLifetime) * time.Second)
	return db, cleanFunc, nil
}

// AutoMigrate 自动映射数据表
func AutoMigrate(db *gorm.DB) error {
	if dbType := config.C.Gorm.DBType; strings.ToLower(dbType) == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}
	return db.AutoMigrate(
		new(entity.Order),
		new(entity.OrderItem),
		new(entity.OrderItemMessage),
	)
}
