package service

import (
	"context"
	"errors"
	"time"

	"github.com/drrrMikado/shorten/internal/domain/alias"
	"github.com/drrrMikado/shorten/pkg/encode"
	"github.com/drrrMikado/shorten/pkg/snowflake"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(New)

type Service struct {
	alias    alias.Usecase
	idworker *snowflake.IDWorker
	log      *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger, idworker *snowflake.IDWorker, alias alias.Usecase) (*Service, error) {
	logger = logger.Named("service")
	return &Service{
		alias:    alias,
		idworker: idworker,
		log:      logger,
	}, nil
}

func (s *Service) Shorten(ctx context.Context, url string, expire time.Time) (*alias.Alias, error) {
	keyID := s.idworker.NextID()
	key := encode.ToBase62(uint64(keyID))
	a := &alias.Alias{
		Key:    key,
		URL:    url,
		Expire: expire,
	}
	a, err := s.alias.Save(ctx, a)
	if err != nil {
		s.log.Errorw("save alias failed", "error", err)
		return nil, err
	}
	return a, nil
}

func (s *Service) Redirect(ctx context.Context, key string) (string, error) {
	a, err := s.alias.Get(ctx, key)
	if err != nil {
		return "", err
	} else if !a.Expire.IsZero() && a.Expire.Before(time.Now()) {
		return "", errors.New("the alias had expired")
	}
	go func() {
		_ = s.alias.IncrPV(context.Background(), a.ID)
	}()
	return a.URL, nil
}
