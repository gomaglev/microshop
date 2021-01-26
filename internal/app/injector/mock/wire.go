// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package mock

import (
	"github.com/gomaglev/microshop/v1/internal/app/injector"
	"github.com/gomaglev/microshop/v1/internal/app/model/gorm/model"
	"github.com/gomaglev/microshop/v1/internal/app/service"
	"github.com/google/wire"
)

// BuildMockInjector
func BuildMockInjector() (*MockInjector, func(), error) {
	wire.Build(
		injector.InitGormDB,
		model.ModelSet,
		service.ServiceSet,
		MockInjectorSet,
	)
	return new(MockInjector), nil, nil
}
