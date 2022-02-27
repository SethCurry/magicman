package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SubType holds the schema definition for the SubType entity.
type SubType struct {
	ent.Schema
}

// Fields of the SubType.
func (SubType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

// Edges of the SubType.
func (SubType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cards", Card.Type),
	}
}
