package database

import (
	"context"
	"log"

	"github.com/drrrMikado/shorten/conf"
	"github.com/go-redis/redis/v8"
)

// NewRedisClient return a redis client.
func NewRedisClient(c conf.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatalln(err)
	}
	return client
}
