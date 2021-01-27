// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package mock

import (
	"github.com/gomaglev/microshop/internal/app/model/gorm"
	"github.com/gomaglev/microshop/internal/app/model/gorm/model"
	"github.com/google/wire"
)

// BuildModelInjector for tests in service
func BuildModelInjector() (*ModelInjector, func(), error) {
	wire.Build(
		gorm.InitGormDB,
		model.ModelSet,
		ModelInjectorSet,
	)
	return new(ModelInjector), nil, nil
}
