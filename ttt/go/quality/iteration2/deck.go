package cards

import "math/rand"

type Deck []Card

func (d Deck) Deal() Card {
	return Card{}
}

func (d Deck) Shuffle() {
	for i, v := range rand.Perm(len(d)) {
		d[v] = d[i]
	}
}

func (d Deck) Contains(c Card) bool {
	return false
}

func (d Deck) Remove(c Card) int {
	return d.RemoveIf(func(current Card) bool {
		return c == current
	})
}

func (d Deck) RemoveIf(fn func(c Card) bool) int {
	return 0
}
