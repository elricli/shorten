package service

import (
	"context"

	"github.com/drrrMikado/shorten/internal/bloomfilter"
	"github.com/drrrMikado/shorten/internal/config"
	"github.com/drrrMikado/shorten/internal/database"
	"github.com/go-redis/redis/v8"
)

// Service struct.
type Service struct {
	c   *config.Config
	rdb *redis.Client
	bf  *bloomfilter.BloomFilter
}

// New service.
func New(ctx context.Context, cfg *config.Config) (*Service, error) {
	redisClient, err := database.NewRedisClient(ctx, cfg.Redis)
	if err != nil {
		return nil, err
	}
	s := &Service{
		c:   cfg,
		rdb: redisClient,
	}
	// set bloom filter
	s.bf, err = bloomfilter.New(cfg.BloomFilter.ExpectedInsertions, cfg.BloomFilter.FPP, cfg.BloomFilter.HashSeed)
	if err != nil {
		return nil, err
	}
	redisResult, err := s.rdb.HGetAll(ctx, redisHashTableKey).Result()
	if err != nil {
		return nil, err
	}
	for k := range redisResult {
		s.bf.Insert([]byte(k))
	}
	return s, nil
}

// Close service.
func (s *Service) Close() error {
	return s.rdb.Close()
}
