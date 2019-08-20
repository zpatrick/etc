package cards

import "math/rand"

type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	return nil
}

func (d *Deck) Shuffle() {
	for i := range d.cards {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}
