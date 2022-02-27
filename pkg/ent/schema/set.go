package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Set holds the schema definition for the Set entity.
type Set struct {
	ent.Schema
}

// Fields of the Set.
func (Set) Fields() []ent.Field {
	return []ent.Field{
		field.String("code"),
		field.String("name"),
		field.String("type"),
		field.String("border"),
		field.String("release_date"),
	}
}

// Edges of the Set.
func (Set) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cards", Card.Type),
	}
}
