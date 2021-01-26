package authentication

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

type Authenticator struct {
	BasicToken string
}

// Basic authenticate func
func (a *Authenticator) Basic() func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		token, err := grpc_auth.AuthFromMD(ctx, "basic")
		if err != nil {
			return nil, err
		}
		if token != a.BasicToken {
			return nil, status.Errorf(codes.Unauthenticated, "authorization failed")
		}
		return ctx, nil
	}
}
