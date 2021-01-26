package server

import (
	"github.com/gomaglev/microshop/internal/pkg/server/rpc"

	"github.com/google/wire"
)

// ServerSet inject server
var ServerSet = wire.NewSet(
	rpc.ServerSet,
	rpc.GatewaySet,
)