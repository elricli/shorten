package shorturl

import (
	"context"

	"github.com/drrrMikado/shorten/pkg/encode"
	"github.com/drrrMikado/shorten/pkg/generator"
)

type service struct {
	idWorker *generator.IDWorker
	repo     Repository
}

func NewService(repo Repository, worker *generator.IDWorker) Service {
	return &service{
		idWorker: worker,
		repo:     repo,
	}
}

func (s *service) Shorten(ctx context.Context, longUrl string) (*ShortUrl, error) {
	nextID, err := s.idWorker.NextID()
	if err != nil {
		return nil, err
	}
	shortUrl := &ShortUrl{
		Key:     encode.ToBase62(uint64(nextID)),
		LongUrl: longUrl,
	}
	err = s.repo.Create(ctx, shortUrl)
	return shortUrl, err
}

func (s *service) Get(ctx context.Context, key string) (*ShortUrl, error) {
	return s.repo.Get(ctx, key)
}
