package shorturl

import (
	"context"
	"log"

	"github.com/drrrMikado/shorten/internal/idworker"
	"github.com/drrrMikado/shorten/pkg/encode"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Shorten(ctx context.Context, longUrl string) (*ShortUrl, error) {
	id, err := idworker.Get()
	if err != nil {
		return nil, err
	}
	shortUrl := &ShortUrl{
		Key: encode.ToBase62(uint64(id)),
		URL: longUrl,
	}
	err = s.repo.Create(ctx, shortUrl)
	return shortUrl, err
}

func (s *service) Redirect(ctx context.Context, key string) (*ShortUrl, error) {
	shortUrl, err := s.repo.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	go func() {
		if err := s.repo.IncrPV(context.TODO(), shortUrl.ID); err != nil {
			log.Printf("s.repo.IncrPV:%d, error:%v\n", shortUrl.ID, err)
		}
	}()
	return shortUrl, nil
}
