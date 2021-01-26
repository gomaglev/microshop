package authorization

import (
	"context"
)

type Authorizer struct{}

// Enforce resources
func (a *Authorizer) Enforce() func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		// Not used in current version
		return ctx, nil
	}
}
