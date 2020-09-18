package service

import (
	"context"
	"log"

	"github.com/drrrMikado/shorten/conf"
	"github.com/drrrMikado/shorten/internal/bloomfilter"
	"github.com/drrrMikado/shorten/internal/database"
	"github.com/go-redis/redis/v8"
)

// Service struct.
type Service struct {
	c   *conf.Config
	rdb *redis.Client
	bf  *bloomfilter.BloomFilter
}

// New service.
func New(c *conf.Config) *Service {
	s := &Service{
		c:   c,
		rdb: database.NewRedisClient(c.Redis),
	}
	// set bloom filter
	s.bf = bloomfilter.New(c.BloomFilter.ExpectedInsertions, c.BloomFilter.FPP)
	result, err := s.rdb.HGetAll(context.Background(), redisHashKey).Result()
	if err != nil {
		log.Fatalln("redis hgetall err:", err)
	}
	for k := range result {
		s.bf.Insert([]byte(k))
	}
	return s
}

// Close service.
func (s *Service) Close() {
	s.rdb.Close()
}
