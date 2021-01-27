package injector

import (
	"github.com/gomaglev/microshop/internal/pkg/config"
	"github.com/gomaglev/microshop/pkg/grpclimit"

	"go.uber.org/ratelimit"
)

// InitGrpcLimiter return new go-grpc Limiter, specified the number of requests you want to limit as a counts per second.
func InitGrpcLimiter() grpclimit.Limiter {
	limit := 1000
	if config.C.GRPC.RateLimitCount > 0 {
		limit = config.C.GRPC.RateLimitCount
	}
	return grpclimit.Limiter{
		Limiter: ratelimit.New(limit),
	}
}
