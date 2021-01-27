package order

import (
	"context"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	mock "github.com/gomaglev/microshop/internal/app/test/mock/model"
	iutil "github.com/gomaglev/microshop/internal/pkg/util"
	"github.com/gomaglev/protos/pkg/proto/common"
	orderproto "github.com/gomaglev/protos/pkg/proto/order"
	"github.com/stretchr/testify/assert"
)

// Test OrderService
func Test(t *testing.T) {
	var firstId = ""
	var secondId = ""
	var thirdId = ""
	ctx := iutil.InitConfig()
	ctx, cancel := context.WithTimeout(ctx, time.Minute*5)
	defer cancel()
	i, _, err := mock.BuildModelInjector()
	orderService := OrderService{
		OrderModel: i.OrderModel,
	}
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
		created, err := orderService.Create(ctx, &CreatOrderRequest{
			Order: &fakeOrder,
		})
		assert.Nil(t, err)
		assert.NotNil(t, created.Id)
		firstId = created.Id

		// second record
		fakeOrder.Id = ""
		created, err = orderService.Create(ctx, &CreatOrderRequest{
			Order: &fakeOrder,
		})
		assert.Nil(t, err)
		assert.NotNil(t, created.Id)
		secondId = created.Id

		// third record
		fakeOrder.Id = ""
		created, err = orderService.Create(ctx, &CreatOrderRequest{
			Order: &fakeOrder,
		})
		assert.Nil(t, err)
		assert.NotNil(t, created.Id)
		thirdId = created.Id
	})
	t.Run("get order", func(t *testing.T) {
		// get first record
		getRes, err := orderService.Get(ctx, &GetOrderRequest{
			Id: firstId,
		})

		assert.Nil(t, err)
		assert.NotNil(t, getRes)
		assert.Equal(t, firstId, getRes.Order.Id)
		assert.Equal(t, "v2", getRes.Order.Status)
	})
	t.Run("list orders", func(t *testing.T) {
		// list second and third record
		listRes, err := orderService.List(ctx, &ListOrdersRequest{
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
		deleted, err := orderService.Delete(ctx, &DeleteOrderRequest{
			Id: firstId,
		})
		assert.Nil(t, err)
		assert.Equal(t, int64(1), deleted.Deleted)

		// delete second and third record
		deleted, err = orderService.Delete(ctx, &DeleteOrderRequest{
			Ids: []string{secondId, thirdId},
		})
		assert.Nil(t, err)
		assert.Equal(t, int64(2), deleted.Deleted)
	})

}
