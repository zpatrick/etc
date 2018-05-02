package cards

import "fmt"

//go:generate stringer -type=Face

type Face int

const (
	Two Face = 2 << iota
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
	Face  Face
}

func (c *Card) String() string {
	return fmt.Sprintf("%s of %s", c.Face, c.Suite)
}
