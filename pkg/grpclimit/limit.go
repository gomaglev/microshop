package grpclimit

import (
	"go.uber.org/ratelimit"
)

type Limiter struct {
	ratelimit.Limiter
}

func (l *Limiter) Limit() bool {
	l.Take()
	return false
}
