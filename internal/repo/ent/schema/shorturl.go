package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ShortUrl holds the schema definition for the ShortUrl entity.
type ShortUrl struct {
	ent.Schema
}

// Fields of the ShortUrl.
func (ShortUrl) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").NotEmpty().Default(""),
		field.String("url").NotEmpty().Default(""),
		field.Uint64("pv").Default(0).Optional(),
		field.Time("expire").Optional(),
		field.Time("create_at").Default(time.Now),
		field.Time("update_at").Default(time.Now),
	}
}

// Edges of the ShortUrl.
func (ShortUrl) Edges() []ent.Edge {
	return nil
}
