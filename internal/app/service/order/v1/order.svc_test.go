package order

import (
	"testing"
)

const (
	configFile = "../../../configs/config.toml"
)

func TestOrderService(t *testing.T) {
	//ctx := test.InitConfig()
	// injector, _, err := mock.BuildMockInjector()
	// if err != nil {
	// 	logrus.Fatal(errors.Wrap(err, "Failed to build mock injector"))
	// }

	// t.Run("create order", func(t *testing.T) {

	// })

	// // create order test
	// // // Create order
	// // func (a *OrderService) Create(ctx context.Context, req *CreatOrderRequest) (*CreatOrderResponse, error) {
	// // 	return nil, nil
	// // }

	// // create pg states fake data
	// state := stateproto.DeviceState{}
	// if err := faker.FakeData(&state); err != nil {
	// 	t.Fatalf("failed to generate state mock data, %v", err)
	// }
	// state.Sensors = make(map[string]string)
	// state.Status = 1
	// state.UpdatedAt = ptypes.TimestampNow()

	// _, err = injector.StateModel.Create(ctx, &state)
	// assert.Nil(t, err)

	// // create mongo rules fake data
	// rule := ruleproto.Rule{}
	// _ = faker.SetRandomMapAndSliceSize(4)
	// if err := faker.FakeData(&rule); err != nil {
	// 	t.Fatalf("failed to generate rule mock data, %v", err)
	// }
	// rule.Status = "ACTIVATED"
	// rule.Enabled = wrapperspb.Bool(true)
	// rule.ApplicationId = state.ApplicationId
	// rule.Id = primitive.NewObjectID().Hex()

	// // https://dev.mysql.com/doc/refman/8.0/en/comparison-operators.html
	// rule.Triggers = []*ruleproto.Trigger{
	// 	{
	// 		Conditions: []*ruleproto.Condition{
	// 			{
	// 				Operator: "interval",
	// 				Value:    "120",
	// 			},
	// 		},
	// 	}}
	// _, err = injector.RuleModel.Create(ctx, &rule)
	// assert.Nil(t, err)

	// time.Sleep(time.Millisecond * 2000)

	// bllOffline := statusbll.DeviceOffline{
	// 	DeviceStatus: statusbll.DeviceStatus{
	// 		RuleModel:  injector.RuleModel,
	// 		StateModel: injector.StateModel,
	// 	},
	// }

	// _, err = bllOffline.Detect(ctx, "", "")
	// assert.Nil(t, err)

	// bllOnline := statusbll.DeviceOnline{
	// 	DeviceStatus: statusbll.DeviceStatus{
	// 		RuleModel:  injector.RuleModel,
	// 		StateModel: injector.StateModel,
	// 	},
	// }

	// _, err = bllOnline.Detect(ctx, "", "")
	// assert.Nil(t, err)
}
