package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// ShortUrl holds the schema definition for the ShortUrl entity.
type ShortUrl struct {
	ent.Schema
}

// Fields of the ShortUrl.
func (ShortUrl) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").NotEmpty().Default(""),
		field.String("short_url").NotEmpty().Default(""),
		field.String("long_url").NotEmpty().Default(""),
		field.Time("create_at").Default(time.Now),
		field.Time("update_at").Default(time.Now),
	}
}

// Edges of the ShortUrl.
func (ShortUrl) Edges() []ent.Edge {
	return nil
}
