package mock

import (
	"github.com/gomaglev/microshop/v1/internal/app/service/order/v1"
	"github.com/google/wire"
)

// MockInjectorSet inject Injector, only for test
var MockInjectorSet = wire.NewSet(wire.Struct(new(MockInjector), "*"))

type MockInjector struct {
	OrderService *order.OrderService
}
