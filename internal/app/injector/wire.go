// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package injector

import (
	gormModel "github.com/gomaglev/microshop/v1/internal/app/model/gorm/model"
	"github.com/gomaglev/microshop/v1/internal/app/service"
	"github.com/gomaglev/microshop/v1/internal/pkg/server"

	"github.com/google/wire"
)

// BuildInjector
func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGormDB,
		InitGrpcLimiter,
		gormModel.ModelSet,
		// mongoModel.ModelSet,
		service.ServiceSet,
		server.ServerSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
