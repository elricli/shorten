package shorturl

import (
	"context"
	"log"

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
		Key: encode.ToBase62(uint64(nextID)),
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
