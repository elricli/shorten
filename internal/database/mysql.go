package database

import (
	"context"
	"fmt"

	"github.com/drrrMikado/shorten/internal/repo/ent"
	"github.com/facebook/ent/dialect"
)

// NewDB .
func NewDB(ctx context.Context, dsn string) (*ent.Client, error) {
	client, err := ent.Open(dialect.MySQL, dsn)
	if err != nil {
		return nil, err
	}
	if err = client.Schema.Create(ctx); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}
	return client, nil
}
