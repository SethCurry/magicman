package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SuperType holds the schema definition for the SuperType entity.
type SuperType struct {
	ent.Schema
}

// Fields of the SuperType.
func (SuperType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

// Edges of the SuperType.
func (SuperType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cards", Card.Type),
	}
}
