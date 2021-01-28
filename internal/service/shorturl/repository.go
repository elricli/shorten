package shorturl

import (
	"context"
	"github.com/drrrMikado/shorten/internal/repo/ent"
	"github.com/drrrMikado/shorten/internal/repo/ent/shorturl"
)

type repository struct {
	c *ent.ShortUrlClient
}

func NewRepository(c *ent.Client) Repository {
	return &repository{
		c: c.ShortUrl,
	}
}

func (r *repository) Create(ctx context.Context, record *ShortUrl) error {
	_, err := r.c.Create().
		SetKey(record.Key).
		SetShortURL(record.ShortUrl).
		SetLongURL(record.LongUrl).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Get(ctx context.Context, key string) (*ShortUrl, error) {
	shortUrl, err := r.c.Query().Where(shorturl.Key(key)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &ShortUrl{
		Key:      shortUrl.Key,
		ShortUrl: shortUrl.ShortURL,
		LongUrl:  shortUrl.LongURL,
	}, nil
}
