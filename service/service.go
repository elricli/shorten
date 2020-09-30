package service

import (
	"context"
	"database/sql"

	"github.com/drrrMikado/shorten/internal/config"
	"github.com/drrrMikado/shorten/internal/database"
	"github.com/go-redis/redis/v8"
)

// Service struct.
type Service struct {
	c   *config.Config
	rdb *redis.Client
	// bf  *bloomfilter.BloomFilter
	db *sql.DB
}

// New service.
func New(ctx context.Context, cfg *config.Config) (*Service, error) {
	redis, err := database.NewRedisClient(ctx, cfg.Redis)
	if err != nil {
		return nil, err
	}
	dsn := cfg.DBConnInfo()
	db, err := database.NewPostgres(ctx, dsn)
	if err != nil {
		return nil, err
	}
	s := &Service{
		c:   cfg,
		rdb: redis,
		db:  db,
	}
	// // set bloom filter
	// s.bf, err = bloomfilter.New(cfg.BloomFilter.ExpectedInsertions, cfg.BloomFilter.FPP, cfg.BloomFilter.HashSeed)
	// if err != nil {
	// 	return nil, err
	// }
	// result, err := s.rdb.HGetAll(context.Background(), redisHashTableKey).Result()
	// if err != nil {
	// 	return nil, err
	// }
	// for k := range result {
	// 	s.bf.Insert([]byte(k))
	// }
	return s, nil
}

// Close service.
func (s *Service) Close() {
	s.rdb.Close()
	s.db.Close()
}
