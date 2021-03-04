package repo

import (
	"context"
	"os"

	"entgo.io/ent/dialect"
	"github.com/drrrMikado/shorten/internal/repo/ent"
	"github.com/drrrMikado/shorten/internal/service/shorturl"
	"github.com/google/wire"

	_ "github.com/lib/pq"
)

var Set = wire.NewSet(
	NewRepository,
	wire.FieldsOf(new(*Repository), "ShortUrl"),
)

type Repository struct {
	client *ent.Client

	ShortUrl shorturl.Repository
}

func NewRepository() (repo *Repository, cf func(), err error) {
	var client *ent.Client
	if client, err = ent.Open(dialect.Postgres, os.Getenv("DATA_SOURCE_NAME")); err != nil {
		return
	}
	cf = func() {
		_ = client.Close()
	}
	if err = client.Schema.Create(context.Background()); err != nil {
		return
	}
	repo = &Repository{
		client:   client,
		ShortUrl: shorturl.NewRepository(client),
	}
	return
}
