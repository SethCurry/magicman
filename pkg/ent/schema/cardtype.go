package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CardType holds the schema definition for the CardType entity.
type CardType struct {
	ent.Schema
}

// Fields of the CardType.
func (CardType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

// Edges of the CardType.
func (CardType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cards", Card.Type),
	}
}
