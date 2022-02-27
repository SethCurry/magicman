package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DeckCard holds the schema definition for the DeckCard entity.
type DeckCard struct {
	ent.Schema
}

// Fields of the DeckCard.
func (DeckCard) Fields() []ent.Field {
	return []ent.Field{
		field.Int("count").Positive(),
	}
}

// Edges of the DeckCard.
func (DeckCard) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("deck", Deck.Type).Ref("cards").Unique(),
		edge.From("card", Card.Type).Ref("deck_cards").Unique(),
	}
}
