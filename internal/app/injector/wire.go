// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package injector

import (
	gorm "github.com/gomaglev/microshop/internal/app/model/gorm"
	gormModel "github.com/gomaglev/microshop/internal/app/model/gorm/model"
	"github.com/gomaglev/microshop/internal/app/service"
	"github.com/gomaglev/microshop/internal/pkg/server"

	"github.com/google/wire"
)

// BuildInjector
func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		// redis.InitCache,
		// mongo.InitMongo,
		// mongoModel.ModelSet,
		gorm.InitGormDB,
		gormModel.ModelSet,
		service.ServiceSet,
		server.ServerSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
