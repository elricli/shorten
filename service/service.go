package service

import (
	"context"
	"database/sql"

	"github.com/drrrMikado/shorten/internal/config"
	"github.com/drrrMikado/shorten/internal/database"
	"github.com/drrrMikado/shorten/internal/generator"
	"github.com/go-redis/redis/v8"
)

// Service struct.
type Service struct {
	c        *config.Config
	rdb      *redis.Client
	db       *sql.DB
	idWorker *generator.IDWorker
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
		c:        cfg,
		rdb:      redis,
		db:       db,
		idWorker: generator.NewIDWorker(1, 1),
	}
	return s, nil
}

// Close service.
func (s *Service) Close() {
	s.rdb.Close()
	s.db.Close()
}
