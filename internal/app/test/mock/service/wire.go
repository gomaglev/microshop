// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package mock

import (
	"github.com/gomaglev/microshop/internal/app/model/gorm"
	"github.com/gomaglev/microshop/internal/app/model/gorm/model"
	"github.com/gomaglev/microshop/internal/app/service"
	"github.com/google/wire"
)

// BuildMockInjector for tests in api
func BuildMockInjector() (*MockInjector, func(), error) {
	wire.Build(
		gorm.InitGormDB,
		model.ModelSet,
		service.ServiceSet,
		MockInjectorSet,
	)
	return new(MockInjector), nil, nil
}
