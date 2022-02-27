package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Color holds the schema definition for the Color entity.
type Color struct {
	ent.Schema
}

// Fields of the Color.
func (Color) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("code").Unique(),
	}
}

// Edges of the Color.
func (Color) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cards", Card.Type),
	}
}
