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
func New(ctx context.Context, cfg *conf.Config) *Service {
	s := &Service{
		c:   cfg,
		rdb: database.NewRedisClient(ctx, cfg.Redis),
	}
	// set bloom filter
	s.bf = bloomfilter.New(cfg.BloomFilter.ExpectedInsertions, cfg.BloomFilter.FPP)
	result, err := s.rdb.HGetAll(ctx, redisHashKey).Result()
	if err != nil {
		log.Fatalln("redis HGetAll err:", err)
	}
	for k := range result {
		s.bf.Insert([]byte(k))
	}
	return s
}

// Close service.
func (s *Service) Close() {
	_ = s.rdb.Close()
}
