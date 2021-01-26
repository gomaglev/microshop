package service

import (
	item "github.com/gomaglev/microshop/v1/internal/app/service/order/item/v1"
	order "github.com/gomaglev/microshop/v1/internal/app/service/order/v1"

	"github.com/google/wire"
)

// ServiceSet api injection
var ServiceSet = wire.NewSet(
	RegisterSet,
	order.OrderSet,
	item.ItemSet,
)
