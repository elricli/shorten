package alias

import (
	"context"

	"github.com/drrrMikado/shorten/internal/repo"
	"github.com/drrrMikado/shorten/internal/repo/ent"
	"github.com/drrrMikado/shorten/internal/repo/ent/alias"
)

type repository struct {
	c *ent.Client
}

func NewRepository(c *repo.Repo) Repository {
	return &repository{
		c: c.Client,
	}
}

func (r *repository) Create(ctx context.Context, a *Alias) error {
	_, err := r.c.Alias.Create().
		SetKey(a.Key).
		SetURL(a.URL).
		SetExpire(a.Expire).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Get(ctx context.Context, key string) (*Alias, error) {
	a, err := r.c.Alias.Query().Where(alias.Key(key)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &Alias{
		ID:     a.ID,
		Key:    a.Key,
		URL:    a.URL,
		Expire: a.Expire,
	}, nil
}

func (r *repository) IncrPV(ctx context.Context, id int) error {
	return r.c.Alias.Update().
		AddPv(1).
		Where(alias.ID(id)).Exec(ctx)
}
