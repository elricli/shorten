package shorturl

import (
	"context"
	"time"

	"github.com/drrrMikado/shorten/internal/repo/ent"
)

type ShortUrl struct {
	ID     int
	Key    string
	URL    string
	Expire time.Time
}

type Service interface {
	Shorten(ctx context.Context, longUrl string, expire time.Time) (*ShortUrl, error)
	Redirect(ctx context.Context, key string) (*ShortUrl, error)
}

type Repository interface {
	Create(ctx context.Context, record *ShortUrl) error
	Get(ctx context.Context, key string) (*ShortUrl, error)
	IncrPV(ctx context.Context, id int) error
}

func FromEnt(shortUrl *ent.ShortUrl) *ShortUrl {
	return &ShortUrl{
		ID:  shortUrl.ID,
		Key: shortUrl.Key,
		URL: shortUrl.URL,
	}
}
