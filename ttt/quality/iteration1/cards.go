package cards

type Card struct {
	Suite string
	Face  string
}

func NewDeck() []*Card {
	suites := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	faces := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	cards := []*Card{}
	for _, suite := range suites {
		for _, face := range faces {
			card := &Card{suite, face}
			cards = append(cards, card)
		}
	}

	return cards
}

func NewShuffledDeck() []*Card {
	deck := NewDeck()

	return nil
}
