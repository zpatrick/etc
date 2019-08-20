package cards

import "fmt"

//go:generate stringer -type=Face

type Rank int

const (
	Two Rank = 2 << iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Suite string

const (
	Clubs    = "Clubs"
	Diamonds = "Diamonds"
	Hearts   = "Hearts"
	Spades   = "Spades"
)

type Card struct {
	Suite Suite
	Rank  Rank
}

func (c *Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank, c.Suite)
}

func (c *Card) IsFace() bool {
	return IsFace(c.Rank)
}

func IsFace(r Rank) bool {
	return r >= Jack
}
