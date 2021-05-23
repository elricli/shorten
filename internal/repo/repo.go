package repo

import (
	"context"

	"github.com/drrrMikado/shorten/internal/repo/ent"
	"github.com/google/wire"

	_ "github.com/lib/pq"
)

var ProviderSet = wire.NewSet(New)

type (
	Repo struct {
		Client *ent.Client
	}
	Config struct {
		Dialect string
		DSN     string
	}
)

func New(cfg Config) (repo *Repo, cf func(), err error) {
	var client *ent.Client
	if client, err = ent.Open(cfg.Dialect, cfg.DSN); err != nil {
		return
	}
	cf = func() {
		_ = client.Close()
	}
	if err = client.Schema.Create(context.Background()); err != nil {
		return
	}
	repo = &Repo{Client: client}
	return
}
