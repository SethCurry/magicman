package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Deck holds the schema definition for the Deck entity.
type Deck struct {
	ent.Schema
}

// Fields of the Deck.
func (Deck) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
	}
}

// Edges of the Deck.
func (Deck) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cards", DeckCard.Type),
	}
}
