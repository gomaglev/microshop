package model

import (
	"context"

	"github.com/gomaglev/microshop/internal/app/model"
	"github.com/gomaglev/microshop/pkg/icontext"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ model.ITrans = new(Trans)

// TransSet
var TransSet = wire.NewSet(wire.Struct(new(Trans), "*"), wire.Bind(new(model.ITrans), new(*Trans)))

// Trans
type Trans struct {
	Client *mongo.Client
}

// Exec
func (a *Trans) Exec(ctx context.Context, fn func(context.Context) error) error {
	if _, ok := icontext.FromTrans(ctx); ok {
		return fn(ctx)
	}

	session, err := a.Client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		err := fn(icontext.NewTrans(sessCtx, true))
		return nil, err
	})

	if err != nil {
		return err
	}
	return nil
}
