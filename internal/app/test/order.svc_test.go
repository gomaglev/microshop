package test

import (
	"testing"

	"github.com/bxcodec/faker"
	"github.com/gomaglev/microshop/internal/app/injector/mock"
	"github.com/gomaglev/microshop/internal/app/service/order/v1"
	iutil "github.com/gomaglev/microshop/internal/pkg/util"
	"github.com/gomaglev/microshop/pkg/logger"
	orderproto "github.com/gomaglev/protos/pkg/proto/order/v1"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestOrderService(t *testing.T) {
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
	res, err := orderService.Create(ctx, &order.CreatOrderRequest{
		Order: &fakeOrder,
	})
	assert.Nil(t, err)
	logger.Printf(ctx, "%v", res)
	assert.NotNil(t, res)
}