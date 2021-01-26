package model

import (
	"context"
)

// ITrans
type ITrans interface {
	// Exec with Transaction
	Exec(ctx context.Context, fn func(context.Context) error) error
}
