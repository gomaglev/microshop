package injector

import (
	"github.com/gomaglev/microshop/internal/pkg/server/rpc"

	"github.com/google/wire"
)

// InjectorSet inject Injector
var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

// Injector injections in app start
type Injector struct {
	Server *rpc.Server
	// Databus    databus.Databus
}
