package icontext

import (
	"context"
)

// Keys
type (
	transCtx     struct{}
	noTransCtx   struct{}
	transLockCtx struct{}
	userIDCtx    struct{}
	traceIDCtx   struct{}
)

// NewTrans create new trans
func NewTrans(ctx context.Context, trans interface{}) context.Context {
	return context.WithValue(ctx, transCtx{}, trans)
}

// FromTrans get from trans
func FromTrans(ctx context.Context) (interface{}, bool) {
	v := ctx.Value(transCtx{})
	return v, v != nil
}

// NewNoTrans create without trans
func NewNoTrans(ctx context.Context) context.Context {
	return context.WithValue(ctx, noTransCtx{}, true)
}

// FromNoTrans from without trans
func FromNoTrans(ctx context.Context) bool {
	v := ctx.Value(noTransCtx{})
	return v != nil && v.(bool)
}

// NewTransLock create trans lock
func NewTransLock(ctx context.Context) context.Context {
	return context.WithValue(ctx, transLockCtx{}, true)
}

// FromTransLock from trans lock
func FromTransLock(ctx context.Context) bool {
	v := ctx.Value(transLockCtx{})
	return v != nil && v.(bool)
}
