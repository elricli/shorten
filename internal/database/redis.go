package database

import (
	"context"
	"log"

	"github.com/drrrMikado/shorten/conf"
	"github.com/go-redis/redis/v8"
)

// NewRedisClient return a redis client.
func NewRedisClient(ctx context.Context, cfg conf.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalln(err)
	}
	return client
}
