// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/SethCurry/magicman/pkg/ent/supertype"
)

// SuperType is the model entity for the SuperType schema.
type SuperType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SuperTypeQuery when eager-loading is set.
	Edges SuperTypeEdges `json:"edges"`
}

// SuperTypeEdges holds the relations/edges for other nodes in the graph.
type SuperTypeEdges struct {
	// Cards holds the value of the cards edge.
	Cards []*Card `json:"cards,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CardsOrErr returns the Cards value or an error if the edge
// was not loaded in eager-loading.
func (e SuperTypeEdges) CardsOrErr() ([]*Card, error) {
	if e.loadedTypes[0] {
		return e.Cards, nil
	}
	return nil, &NotLoadedError{edge: "cards"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SuperType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case supertype.FieldID:
			values[i] = new(sql.NullInt64)
		case supertype.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SuperType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SuperType fields.
func (st *SuperType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case supertype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			st.ID = int(value.Int64)
		case supertype.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				st.Name = value.String
			}
		}
	}
	return nil
}

// QueryCards queries the "cards" edge of the SuperType entity.
func (st *SuperType) QueryCards() *CardQuery {
	return (&SuperTypeClient{config: st.config}).QueryCards(st)
}

// Update returns a builder for updating this SuperType.
// Note that you need to call SuperType.Unwrap() before calling this method if this SuperType
// was returned from a transaction, and the transaction was committed or rolled back.
func (st *SuperType) Update() *SuperTypeUpdateOne {
	return (&SuperTypeClient{config: st.config}).UpdateOne(st)
}

// Unwrap unwraps the SuperType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (st *SuperType) Unwrap() *SuperType {
	tx, ok := st.config.driver.(*txDriver)
	if !ok {
		panic("ent: SuperType is not a transactional entity")
	}
	st.config.driver = tx.drv
	return st
}

// String implements the fmt.Stringer.
func (st *SuperType) String() string {
	var builder strings.Builder
	builder.WriteString("SuperType(")
	builder.WriteString(fmt.Sprintf("id=%v", st.ID))
	builder.WriteString(", name=")
	builder.WriteString(st.Name)
	builder.WriteByte(')')
	return builder.String()
}

// SuperTypes is a parsable slice of SuperType.
type SuperTypes []*SuperType

func (st SuperTypes) config(cfg config) {
	for _i := range st {
		st[_i].config = cfg
	}
}
