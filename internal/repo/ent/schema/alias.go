package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Alias holds the schema definition for the Alias entity.
type Alias struct {
	ent.Schema
}

// Fields of the Alias.
func (Alias) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").NotEmpty().Default(""),
		field.String("url").NotEmpty().Default(""),
		field.Uint64("pv").Default(0).Optional(),
		field.Time("expire").Optional(),
	}
}

// Edges of the Alias.
func (Alias) Edges() []ent.Edge {
	return nil
}

func (Alias) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
