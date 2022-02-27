// Code generated by entc, DO NOT EDIT.

package supertype

const (
	// Label holds the string label denoting the supertype type in the database.
	Label = "super_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeCards holds the string denoting the cards edge name in mutations.
	EdgeCards = "cards"
	// Table holds the table name of the supertype in the database.
	Table = "super_types"
	// CardsTable is the table that holds the cards relation/edge. The primary key declared below.
	CardsTable = "super_type_cards"
	// CardsInverseTable is the table name for the Card entity.
	// It exists in this package in order to avoid circular dependency with the "card" package.
	CardsInverseTable = "cards"
)

// Columns holds all SQL columns for supertype fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// CardsPrimaryKey and CardsColumn2 are the table columns denoting the
	// primary key for the cards relation (M2M).
	CardsPrimaryKey = []string{"super_type_id", "card_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
