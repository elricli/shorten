package redis

import (
	"github.com/drrrMikado/shorten/conf"
	"github.com/go-redis/redis/v8"
)

// Client redis client.
var Client *redis.Client

// Init return a redis client.
func Init(c conf.RedisConf) {
	Client = redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})
}
