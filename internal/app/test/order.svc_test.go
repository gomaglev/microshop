package test

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/gomaglev/microshop/internal/app/injector/mock"
	"github.com/gomaglev/microshop/internal/app/service/v1/order"
	iutil "github.com/gomaglev/microshop/internal/pkg/util"
	orderproto "github.com/gomaglev/protos/pkg/proto/order"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestOrderService(t *testing.T) {
	t.Run("create order", func(t *testing.T) {
		ctx := iutil.InitConfig()
		injector, _, err := mock.BuildMockInjector()
		if err != nil {
			logrus.Fatal(errors.Wrap(err, "Failed to build mock injector"))
		}
		orderService := *injector.OrderService
		fakeOrder := orderproto.Order{}
		if err := faker.FakeData(&fakeOrder); err != nil {
			t.Fatalf("failed to generate order mock data, %v", err)
		}
		fakeOrder.Id = ""
		res, err := orderService.Create(ctx, &order.CreatOrderRequest{
			Order: &fakeOrder,
		})
		assert.Nil(t, err)
		assert.NotNil(t, res.Id)
	})
}
