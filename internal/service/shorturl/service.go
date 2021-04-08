package shorturl

import (
	"context"
	"errors"
	"time"

	"github.com/drrrMikado/shorten/internal/idworker"
	"github.com/drrrMikado/shorten/pkg/encode"
	"github.com/drrrMikado/shorten/pkg/log"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Shorten(ctx context.Context, longUrl string, expire time.Time) (*ShortUrl, error) {
	id, err := idworker.Get()
	if err != nil {
		log.Errorw("Failed get id by id worker", "err", err)
		return nil, err
	}
	shortUrl := &ShortUrl{
		Key: encode.ToBase62(uint64(id)),
		URL: longUrl,
	}
	if !expire.IsZero() {
		shortUrl.Expire = expire
	}
	err = s.repo.Create(ctx, shortUrl)
	if err != nil {
		log.Errorw("Failed get by key",
			"key", shortUrl.Key,
			"url", shortUrl.URL,
			"err", err,
		)
		return nil, err
	}
	return shortUrl, nil
}

func (s *service) Redirect(ctx context.Context, key string) (*ShortUrl, error) {
	shortUrl, err := s.repo.Get(ctx, key)
	if err != nil {
		log.Errorw("Failed get by key",
			"key", key,
			"err", err,
		)
		return nil, err
	}
	if !shortUrl.Expire.IsZero() && shortUrl.Expire.Before(time.Now()) {
		return nil, errors.New("not found")
	}
	go func() {
		if err := s.repo.IncrPV(context.TODO(), shortUrl.ID); err != nil {
			log.Errorw("increase pv error",
				"id", shortUrl.ID,
				"err", err,
			)
		}
	}()
	return shortUrl, nil
}
