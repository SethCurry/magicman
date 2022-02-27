package magicman

type DeckList struct {
	Commander string
	Cards     []DeckListCard
}

type DeckListCard struct {
	Name  string
	Count int
}
