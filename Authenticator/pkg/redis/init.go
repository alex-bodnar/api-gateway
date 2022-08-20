package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"

	"authenticator/pkg/log"
)

// InitRedis opens connection and ping redis
func InitRedis(ctx context.Context, cfg Config, logger log.Logger) *redis.Client {
	logger.Infof("open redis: connection to %q", cfg.ConnectionAddr)

	rdb, err := open(ctx, cfg)
	if err == nil {
		return rdb
	}

	for i := 0; i < cfg.MaxRetries; i++ {
		rdb, err := open(ctx, cfg)
		if err != nil {
			logger.Errorf("retry: connection to %q: %v", cfg.ConnectionAddr, err)
			time.Sleep(cfg.RetryDelay)
			continue
		}

		return rdb
	}

	logger.Fatalf("failed: open connection to %q: %w\n", cfg.ConnectionAddr, err)

	return nil
}

func open(ctx context.Context, cfg Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.ConnectionAddr,
		Password: cfg.Pass,
		DB:       cfg.DB,
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return rdb, nil
}
