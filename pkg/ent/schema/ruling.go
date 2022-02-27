package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Ruling holds the schema definition for the Ruling entity.
type Ruling struct {
	ent.Schema
}

// Fields of the Ruling.
func (Ruling) Fields() []ent.Field {
	return []ent.Field{
		field.String("text"),
		field.Time("date"),
	}
}

// Edges of the Ruling.
func (Ruling) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("card", Card.Type).Ref("rulings").Unique(),
	}
}
