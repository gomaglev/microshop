package test

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/gomaglev/microshop/internal/app/injector/mock"
	orderv1 "github.com/gomaglev/microshop/internal/app/service/v1/order"
	iutil "github.com/gomaglev/microshop/internal/pkg/util"
	"github.com/gomaglev/protos/pkg/proto/common"
	orderproto "github.com/gomaglev/protos/pkg/proto/order"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestOrderService(t *testing.T) {
	var firstId = ""
	var secondId = ""
	var thirdId = ""
	ctx := iutil.InitConfig()
	injector, _, err := mock.BuildMockInjector()
	orderServiceV1 := *injector.OrderService
	if err != nil {
		logrus.Fatal(errors.Wrap(err, "failed to build mock injector"))
	}
	t.Run("create order", func(t *testing.T) {
		fakeOrder := orderproto.Order{}
		if err := faker.FakeData(&fakeOrder); err != nil {
			t.Fatalf("failed to generate order mock data, %v", err)
		}
		// first record
		fakeOrder.Id = ""
		created, err := orderServiceV1.Create(ctx, &orderv1.CreatOrderRequest{
			Order: &fakeOrder,
		})
		assert.Nil(t, err)
		assert.NotNil(t, created.Id)
		firstId = created.Id

		// second record
		fakeOrder.Id = ""
		created, err = orderServiceV1.Create(ctx, &orderv1.CreatOrderRequest{
			Order: &fakeOrder,
		})
		assert.Nil(t, err)
		assert.NotNil(t, created.Id)
		secondId = created.Id

		// third record
		fakeOrder.Id = ""
		created, err = orderServiceV1.Create(ctx, &orderv1.CreatOrderRequest{
			Order: &fakeOrder,
		})
		assert.Nil(t, err)
		assert.NotNil(t, created.Id)
		thirdId = created.Id
	})
	t.Run("get order", func(t *testing.T) {
		// get first record
		orderServiceV1 := *injector.OrderService
		getRes, err := orderServiceV1.Get(ctx, &orderv1.GetOrderRequest{
			Id: firstId,
		})

		assert.Nil(t, err)
		assert.NotNil(t, getRes)
		assert.Equal(t, firstId, getRes.Order.Id)
	})
	t.Run("list orders", func(t *testing.T) {
		// list second and third record
		orderServiceV1 := *injector.OrderService
		listRes, err := orderServiceV1.List(ctx, &orderv1.ListOrdersRequest{
			Ids: []string{secondId, thirdId},
			Pagination: &common.PaginationParam{
				Pagination: true,
				Page:       1,
				PageSize:   10,
			},
		})

		assert.Nil(t, err)
		assert.Equal(t, int32(1), listRes.Orders.Pagination.Page)
		assert.Equal(t, 2, len(listRes.Orders.List))
	})

	t.Run("delete orders", func(t *testing.T) {
		// delete first record
		deleted, err := orderServiceV1.Delete(ctx, &orderv1.DeleteOrderRequest{
			Id: firstId,
		})
		assert.Nil(t, err)
		assert.Equal(t, int64(1), deleted.Deleted)

		// delete second and third record
		deleted, err = orderServiceV1.Delete(ctx, &orderv1.DeleteOrderRequest{
			Ids: []string{secondId, thirdId},
		})
		assert.Nil(t, err)
		assert.Equal(t, int64(2), deleted.Deleted)
	})
}
