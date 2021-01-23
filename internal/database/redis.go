package database

import (
	"context"

	"github.com/drrrMikado/shorten/internal/config"
	"github.com/go-redis/redis/v8"
)

// NewRedis return a redis client.
func NewRedis(ctx context.Context, cfg config.Redis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return client, nil
}
