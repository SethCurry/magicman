// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/SethCurry/magicman/pkg/ent/set"
)

// Set is the model entity for the Set schema.
type Set struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Border holds the value of the "border" field.
	Border string `json:"border,omitempty"`
	// ReleaseDate holds the value of the "release_date" field.
	ReleaseDate string `json:"release_date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SetQuery when eager-loading is set.
	Edges SetEdges `json:"edges"`
}

// SetEdges holds the relations/edges for other nodes in the graph.
type SetEdges struct {
	// Cards holds the value of the cards edge.
	Cards []*Card `json:"cards,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CardsOrErr returns the Cards value or an error if the edge
// was not loaded in eager-loading.
func (e SetEdges) CardsOrErr() ([]*Card, error) {
	if e.loadedTypes[0] {
		return e.Cards, nil
	}
	return nil, &NotLoadedError{edge: "cards"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Set) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case set.FieldID:
			values[i] = new(sql.NullInt64)
		case set.FieldCode, set.FieldName, set.FieldType, set.FieldBorder, set.FieldReleaseDate:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Set", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Set fields.
func (s *Set) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case set.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case set.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				s.Code = value.String
			}
		case set.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case set.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				s.Type = value.String
			}
		case set.FieldBorder:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field border", values[i])
			} else if value.Valid {
				s.Border = value.String
			}
		case set.FieldReleaseDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field release_date", values[i])
			} else if value.Valid {
				s.ReleaseDate = value.String
			}
		}
	}
	return nil
}

// QueryCards queries the "cards" edge of the Set entity.
func (s *Set) QueryCards() *CardQuery {
	return (&SetClient{config: s.config}).QueryCards(s)
}

// Update returns a builder for updating this Set.
// Note that you need to call Set.Unwrap() before calling this method if this Set
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Set) Update() *SetUpdateOne {
	return (&SetClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Set entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Set) Unwrap() *Set {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Set is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Set) String() string {
	var builder strings.Builder
	builder.WriteString("Set(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", code=")
	builder.WriteString(s.Code)
	builder.WriteString(", name=")
	builder.WriteString(s.Name)
	builder.WriteString(", type=")
	builder.WriteString(s.Type)
	builder.WriteString(", border=")
	builder.WriteString(s.Border)
	builder.WriteString(", release_date=")
	builder.WriteString(s.ReleaseDate)
	builder.WriteByte(')')
	return builder.String()
}

// Sets is a parsable slice of Set.
type Sets []*Set

func (s Sets) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
