// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CardsColumns holds the columns for the "cards" table.
	CardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "multiverse_id", Type: field.TypeString},
		{Name: "gatherer_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "text", Type: field.TypeString},
		{Name: "cmc", Type: field.TypeInt},
		{Name: "mana_cost", Type: field.TypeString},
		{Name: "artist", Type: field.TypeString},
		{Name: "power", Type: field.TypeString},
		{Name: "toughness", Type: field.TypeString},
		{Name: "image_url", Type: field.TypeString},
		{Name: "original_text", Type: field.TypeString},
		{Name: "original_type", Type: field.TypeString},
		{Name: "cached_image_path", Type: field.TypeString, Nullable: true},
		{Name: "rarity", Type: field.TypeString},
		{Name: "set_cards", Type: field.TypeInt, Nullable: true},
	}
	// CardsTable holds the schema information for the "cards" table.
	CardsTable = &schema.Table{
		Name:       "cards",
		Columns:    CardsColumns,
		PrimaryKey: []*schema.Column{CardsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cards_sets_cards",
				Columns:    []*schema.Column{CardsColumns[16]},
				RefColumns: []*schema.Column{SetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CardTypesColumns holds the columns for the "card_types" table.
	CardTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// CardTypesTable holds the schema information for the "card_types" table.
	CardTypesTable = &schema.Table{
		Name:       "card_types",
		Columns:    CardTypesColumns,
		PrimaryKey: []*schema.Column{CardTypesColumns[0]},
	}
	// ColorsColumns holds the columns for the "colors" table.
	ColorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "code", Type: field.TypeString, Unique: true},
	}
	// ColorsTable holds the schema information for the "colors" table.
	ColorsTable = &schema.Table{
		Name:       "colors",
		Columns:    ColorsColumns,
		PrimaryKey: []*schema.Column{ColorsColumns[0]},
	}
	// DecksColumns holds the columns for the "decks" table.
	DecksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
	}
	// DecksTable holds the schema information for the "decks" table.
	DecksTable = &schema.Table{
		Name:       "decks",
		Columns:    DecksColumns,
		PrimaryKey: []*schema.Column{DecksColumns[0]},
	}
	// DeckCardsColumns holds the columns for the "deck_cards" table.
	DeckCardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "count", Type: field.TypeInt},
		{Name: "card_deck_cards", Type: field.TypeInt, Nullable: true},
		{Name: "deck_cards", Type: field.TypeInt, Nullable: true},
	}
	// DeckCardsTable holds the schema information for the "deck_cards" table.
	DeckCardsTable = &schema.Table{
		Name:       "deck_cards",
		Columns:    DeckCardsColumns,
		PrimaryKey: []*schema.Column{DeckCardsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "deck_cards_cards_deck_cards",
				Columns:    []*schema.Column{DeckCardsColumns[2]},
				RefColumns: []*schema.Column{CardsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "deck_cards_decks_cards",
				Columns:    []*schema.Column{DeckCardsColumns[3]},
				RefColumns: []*schema.Column{DecksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RulingsColumns holds the columns for the "rulings" table.
	RulingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "text", Type: field.TypeString},
		{Name: "date", Type: field.TypeTime},
		{Name: "card_rulings", Type: field.TypeInt, Nullable: true},
	}
	// RulingsTable holds the schema information for the "rulings" table.
	RulingsTable = &schema.Table{
		Name:       "rulings",
		Columns:    RulingsColumns,
		PrimaryKey: []*schema.Column{RulingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "rulings_cards_rulings",
				Columns:    []*schema.Column{RulingsColumns[3]},
				RefColumns: []*schema.Column{CardsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SetsColumns holds the columns for the "sets" table.
	SetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "code", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "border", Type: field.TypeString},
		{Name: "release_date", Type: field.TypeString},
	}
	// SetsTable holds the schema information for the "sets" table.
	SetsTable = &schema.Table{
		Name:       "sets",
		Columns:    SetsColumns,
		PrimaryKey: []*schema.Column{SetsColumns[0]},
	}
	// SubTypesColumns holds the columns for the "sub_types" table.
	SubTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// SubTypesTable holds the schema information for the "sub_types" table.
	SubTypesTable = &schema.Table{
		Name:       "sub_types",
		Columns:    SubTypesColumns,
		PrimaryKey: []*schema.Column{SubTypesColumns[0]},
	}
	// SuperTypesColumns holds the columns for the "super_types" table.
	SuperTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// SuperTypesTable holds the schema information for the "super_types" table.
	SuperTypesTable = &schema.Table{
		Name:       "super_types",
		Columns:    SuperTypesColumns,
		PrimaryKey: []*schema.Column{SuperTypesColumns[0]},
	}
	// CardTypeCardsColumns holds the columns for the "card_type_cards" table.
	CardTypeCardsColumns = []*schema.Column{
		{Name: "card_type_id", Type: field.TypeInt},
		{Name: "card_id", Type: field.TypeInt},
	}
	// CardTypeCardsTable holds the schema information for the "card_type_cards" table.
	CardTypeCardsTable = &schema.Table{
		Name:       "card_type_cards",
		Columns:    CardTypeCardsColumns,
		PrimaryKey: []*schema.Column{CardTypeCardsColumns[0], CardTypeCardsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "card_type_cards_card_type_id",
				Columns:    []*schema.Column{CardTypeCardsColumns[0]},
				RefColumns: []*schema.Column{CardTypesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "card_type_cards_card_id",
				Columns:    []*schema.Column{CardTypeCardsColumns[1]},
				RefColumns: []*schema.Column{CardsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ColorCardsColumns holds the columns for the "color_cards" table.
	ColorCardsColumns = []*schema.Column{
		{Name: "color_id", Type: field.TypeInt},
		{Name: "card_id", Type: field.TypeInt},
	}
	// ColorCardsTable holds the schema information for the "color_cards" table.
	ColorCardsTable = &schema.Table{
		Name:       "color_cards",
		Columns:    ColorCardsColumns,
		PrimaryKey: []*schema.Column{ColorCardsColumns[0], ColorCardsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "color_cards_color_id",
				Columns:    []*schema.Column{ColorCardsColumns[0]},
				RefColumns: []*schema.Column{ColorsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "color_cards_card_id",
				Columns:    []*schema.Column{ColorCardsColumns[1]},
				RefColumns: []*schema.Column{CardsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// SubTypeCardsColumns holds the columns for the "sub_type_cards" table.
	SubTypeCardsColumns = []*schema.Column{
		{Name: "sub_type_id", Type: field.TypeInt},
		{Name: "card_id", Type: field.TypeInt},
	}
	// SubTypeCardsTable holds the schema information for the "sub_type_cards" table.
	SubTypeCardsTable = &schema.Table{
		Name:       "sub_type_cards",
		Columns:    SubTypeCardsColumns,
		PrimaryKey: []*schema.Column{SubTypeCardsColumns[0], SubTypeCardsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sub_type_cards_sub_type_id",
				Columns:    []*schema.Column{SubTypeCardsColumns[0]},
				RefColumns: []*schema.Column{SubTypesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "sub_type_cards_card_id",
				Columns:    []*schema.Column{SubTypeCardsColumns[1]},
				RefColumns: []*schema.Column{CardsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// SuperTypeCardsColumns holds the columns for the "super_type_cards" table.
	SuperTypeCardsColumns = []*schema.Column{
		{Name: "super_type_id", Type: field.TypeInt},
		{Name: "card_id", Type: field.TypeInt},
	}
	// SuperTypeCardsTable holds the schema information for the "super_type_cards" table.
	SuperTypeCardsTable = &schema.Table{
		Name:       "super_type_cards",
		Columns:    SuperTypeCardsColumns,
		PrimaryKey: []*schema.Column{SuperTypeCardsColumns[0], SuperTypeCardsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "super_type_cards_super_type_id",
				Columns:    []*schema.Column{SuperTypeCardsColumns[0]},
				RefColumns: []*schema.Column{SuperTypesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "super_type_cards_card_id",
				Columns:    []*schema.Column{SuperTypeCardsColumns[1]},
				RefColumns: []*schema.Column{CardsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CardsTable,
		CardTypesTable,
		ColorsTable,
		DecksTable,
		DeckCardsTable,
		RulingsTable,
		SetsTable,
		SubTypesTable,
		SuperTypesTable,
		CardTypeCardsTable,
		ColorCardsTable,
		SubTypeCardsTable,
		SuperTypeCardsTable,
	}
)

func init() {
	CardsTable.ForeignKeys[0].RefTable = SetsTable
	DeckCardsTable.ForeignKeys[0].RefTable = CardsTable
	DeckCardsTable.ForeignKeys[1].RefTable = DecksTable
	RulingsTable.ForeignKeys[0].RefTable = CardsTable
	CardTypeCardsTable.ForeignKeys[0].RefTable = CardTypesTable
	CardTypeCardsTable.ForeignKeys[1].RefTable = CardsTable
	ColorCardsTable.ForeignKeys[0].RefTable = ColorsTable
	ColorCardsTable.ForeignKeys[1].RefTable = CardsTable
	SubTypeCardsTable.ForeignKeys[0].RefTable = SubTypesTable
	SubTypeCardsTable.ForeignKeys[1].RefTable = CardsTable
	SuperTypeCardsTable.ForeignKeys[0].RefTable = SuperTypesTable
	SuperTypeCardsTable.ForeignKeys[1].RefTable = CardsTable
}
