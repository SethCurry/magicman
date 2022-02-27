package mtg

type Card struct {
	Name     string   `json:"name"`
	Names    []string `json:"names"`
	ManaCost string   `json:"manaCost"`
	// TODO should convert to an int
	CMC           float32  `json:"cmc"`
	Colors        []Color  `json:"colors"`
	ColorIdentity []Color  `json:"colorIdentity"`
	Type          string   `json:"type"`
	Supertypes    []string `json:"supertypes"`
	Types         []string `json:"types"`
	Subtypes      []string `json:"subtypes"`
	Rarity        string   `json:"rarity"`
	Set           string   `json:"set"`
	Text          string   `json:"text"`
	Artist        string   `json:"artist"`
	Number        string   `json:"number"`
	// TODO how to handle power and toughness when */* or similar
	Power     string `json:"power"`
	Toughness string `json:"toughness"`
	// TODO should be enum
	Layout string `json:"layout"`
	// TODO could maybe be an int?
	MultiverseID string   `json:"multiverseid"`
	ImageURL     string   `json:"imageUrl"`
	Rulings      []Ruling `json:"rulings"`
	Printings    []string `json:"printings"`
	OriginalText string   `json:"originalText"`
	OriginalType string   `json:"originalType"`
	ID           string   `json:"id"`
}
