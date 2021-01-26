// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package mock

import (
	"github.com/gomaglev/microshop/internal/app/injector"
	gormModel "github.com/gomaglev/microshop/internal/app/model/gorm/model"
	"github.com/gomaglev/microshop/internal/app/service"
	"github.com/gomaglev/microshop/internal/pkg/server"

	"github.com/google/wire"
)

// BuildMockInjector
func BuildMockInjector() (*MockInjector, func(), error) {
	wire.Build(
		injector.InitGormDB,
		gormModel.ModelSet,
		server.ServerSet,
		service.ServiceSet,
		MockInjectorSet,
	)
	return new(MockInjector), nil, nil
}
