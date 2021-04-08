package shorturl

import (
	"context"

	"github.com/drrrMikado/shorten/internal/repo/ent"
	"github.com/drrrMikado/shorten/internal/repo/ent/shorturl"
)

type repository struct {
	c *ent.Client
}

func NewRepository(c *ent.Client) Repository {
	return &repository{
		c: c,
	}
}

func (r *repository) Create(ctx context.Context, shortUrl *ShortUrl) error {
	_, err := r.c.ShortUrl.Create().
		SetKey(shortUrl.Key).
		SetURL(shortUrl.URL).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Get(ctx context.Context, key string) (*ShortUrl, error) {
	shortUrl, err := r.c.ShortUrl.Query().Where(shorturl.Key(key)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return FromEnt(shortUrl), nil
}

func (r *repository) IncrPV(ctx context.Context, id int) error {
	return r.c.ShortUrl.Update().
		AddPv(1).
		Where(shorturl.ID(id)).Exec(ctx)
}
