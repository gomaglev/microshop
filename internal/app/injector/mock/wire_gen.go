// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package mock

import (
	"github.com/gomaglev/microshop/internal/app/injector"
	"github.com/gomaglev/microshop/internal/app/model/gorm/model"
	"github.com/gomaglev/microshop/internal/app/service/order/v1"
)

// Injectors from wire.go:

func BuildMockInjector() (*MockInjector, func(), error) {
	db, cleanup, err := injector.InitGormDB()
	if err != nil {
		return nil, nil, err
	}
	modelOrder := &model.Order{
		DB: db,
	}
	orderService := order.OrderService{
		OrderModel: modelOrder,
	}
	mockInjector := &MockInjector{
		OrderService: &orderService,
	}
	return mockInjector, func() {
		cleanup()
	}, nil
}
