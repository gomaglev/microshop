package model

import (
	"context"

	"github.com/gomaglev/microshop/v1/internal/app/model"
	"github.com/gomaglev/microshop/v1/pkg/icontext"

	"gorm.io/gorm"

	"github.com/google/wire"
	"github.com/pkg/errors"
)

var _ model.ITrans = new(Trans)

// TransSet
var TransSet = wire.NewSet(wire.Struct(new(Trans), "*"), wire.Bind(new(model.ITrans), new(*Trans)))

// Trans
type Trans struct {
	DB *gorm.DB
}

// Exec
func (a *Trans) Exec(ctx context.Context, fn func(context.Context) error) error {
	if _, ok := icontext.FromTrans(ctx); ok {
		return fn(ctx)
	}

	err := (*gorm.DB)(a.DB).Transaction(func(db *gorm.DB) error {
		return fn(icontext.NewTrans(ctx, db))
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
