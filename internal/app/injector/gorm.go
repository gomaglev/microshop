package injector

import (
	"github.com/gomaglev/microshop/v1/internal/pkg/config"

	igorm "github.com/gomaglev/microshop/v1/internal/app/model/gorm"

	"gorm.io/gorm"
)

// InitGormDB 初始化gorm存储
func InitGormDB() (*gorm.DB, func(), error) {
	cfg := config.C.Gorm
	db, cleanFunc, err := igorm.NewDB()
	if err != nil {
		return nil, cleanFunc, err
	}

	if cfg.EnableAutoMigrate {
		err = igorm.AutoMigrate(db)
		if err != nil {
			return nil, cleanFunc, err
		}
	}

	return db, cleanFunc, nil
}
