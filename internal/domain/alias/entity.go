package alias

import (
	"context"
	"time"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewRepository, NewUseCase)

type Alias struct {
	ID     int
	Key    string
	URL    string    `validate:"url"`
	Expire time.Time `validate:"gt"`
}

type Usecase interface {
	Save(ctx context.Context, a *Alias) (*Alias, error)
	Get(ctx context.Context, key string) (*Alias, error)
	IncrPV(ctx context.Context, id int) error
}

type Repository interface {
	Create(ctx context.Context, record *Alias) error
	Get(ctx context.Context, key string) (*Alias, error)
	IncrPV(ctx context.Context, id int) error
}
