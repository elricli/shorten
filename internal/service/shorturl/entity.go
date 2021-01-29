package shorturl

import (
	"context"
)

type ShortUrl struct {
	Key     string
	LongUrl string
}

type Service interface {
	Shorten(ctx context.Context, longUrl string) (*ShortUrl, error)
	Get(ctx context.Context, key string) (*ShortUrl, error)
}

type Repository interface {
	Create(ctx context.Context, record *ShortUrl) error
	Get(ctx context.Context, key string) (*ShortUrl, error)
}
