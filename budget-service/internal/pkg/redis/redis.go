package redis

import (
	"budget-service/internal/pkg/load"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(cfg load.Config) (*redis.Client, error) {
	target := fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr: target,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}
