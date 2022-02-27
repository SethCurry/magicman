package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		// TODO add rulings
		// TODO add card colors
		// TODO add CMC
		field.String("multiverse_id"),
		field.String("gatherer_id"),
		field.String("name"),
		field.String("type"),
		field.String("text"),
		field.Int("cmc"),
		field.String("mana_cost"),
		field.String("artist"),
		field.String("power"),
		field.String("toughness"),
		field.String("image_url"),
		field.String("original_text"),
		field.String("original_type"),
		field.String("cached_image_path").Optional().Nillable(),
		field.String("rarity"),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rulings", Ruling.Type),
		edge.From("set", Set.Type).Ref("cards").Unique(),
		edge.From("types", CardType.Type).Ref("cards"),
		edge.From("subtypes", SubType.Type).Ref("cards"),
		edge.From("supertypes", SuperType.Type).Ref("cards"),
		edge.From("colors", Color.Type).Ref("cards"),
		edge.To("deck_cards", DeckCard.Type),
	}
}
