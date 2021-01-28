package service

import (
	"context"

	"github.com/drrrMikado/shorten/internal/config"
	"github.com/drrrMikado/shorten/internal/database"
	"github.com/drrrMikado/shorten/internal/repo/ent"
	"github.com/drrrMikado/shorten/pkg/generator"
	"github.com/go-redis/redis/v8"
)

// Service struct.
type Service struct {
	c        *config.Config
	rdb      *redis.Client
	db       *ent.Client
	sdb      *ent.ShortUrlClient
	idWorker *generator.IDWorker
}

// New service.
func New(ctx context.Context, cfg *config.Config) (*Service, error) {
	redisClient, err := database.NewRedis(ctx, cfg.Redis)
	if err != nil {
		return nil, err
	}
	entCli, err := database.NewDB(ctx, cfg.DBConnInfo())
	if err != nil {
		return nil, err
	}
	s := &Service{
		c:        cfg,
		rdb:      redisClient,
		db:       entCli,
		sdb:      entCli.ShortUrl,
		idWorker: generator.NewIDWorker(1, 1),
	}
	return s, nil
}

// Close service.
func (s *Service) Close() {
	_ = s.rdb.Close()
	_ = s.db.Close()
}
