package mock

import (
	service "github.com/gomaglev/microshop/internal/app/service/order/v1"

	"github.com/google/wire"
)

// MockInjectorSet inject Injector, only for test
var MockInjectorSet = wire.NewSet(wire.Struct(new(MockInjector), "*"))

type MockInjector struct {
	OrderService service.OrderServiceServer
}
