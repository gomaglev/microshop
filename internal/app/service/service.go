package service

import (
	order "github.com/gomaglev/microshop/internal/app/service/order"
	item "github.com/gomaglev/microshop/internal/app/service/order/item"

	"github.com/google/wire"
)

// ServiceSet api injection
var ServiceSet = wire.NewSet(
	RegisterSet,
	order.OrderSet,
	item.ItemSet,
)
