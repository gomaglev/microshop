package redis

import (
	"context"

	"github.com/gomaglev/microshop/internal/pkg/config"
	"github.com/gomaglev/microshop/pkg/logger"

	"github.com/go-redis/redis/v8"
)

// InitCache create redis pool
func InitCache() (*redis.Client, func(), error) {
	var (
		ctx    = context.Background()
		cancel context.CancelFunc
	)

	cfg := config.C.Redis

	if t := cfg.DialTimeout; t > 0 {
		ctx, cancel = context.WithTimeout(ctx, t)
		defer cancel()
	}

	logger.Infof(ctx, "initializing cache")
	options := &redis.Options{
		Addr:        cfg.Host,
		Password:    cfg.Auth,
		DialTimeout: cfg.DialTimeout,
		DB:          0,
	}
	client := redis.NewClient(options)

	pong, err := client.Ping(ctx).Result()
	logger.Infof(ctx, "redis connection established, ping: %s", pong)

	cleanFunc := func() {
		err := client.Close()
		if err != nil {
			logger.Errorf(ctx, "Redis pool close error: %s", err.Error())
		}
	}

	if err != nil {
		return nil, cleanFunc, err
	}
	return client, cleanFunc, nil
}
