package mock

import (
	"github.com/gomaglev/microshop/internal/app/service/v1/order"
	"github.com/google/wire"
)

// MockInjectorSet inject Injector, only for test
var MockInjectorSet = wire.NewSet(wire.Struct(new(MockInjector), "*"))

type MockInjector struct {
	OrderService *order.OrderService
}
