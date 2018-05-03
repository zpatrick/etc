package cards

import "math/rand"

/*
Card objects
------------------------
We know we need a notion of a card. Since cards just have a suit and rank, it makes sense to store that in a struct.
*/

/*
Suits and Ranks
------------------------
We know there are 4 suits: clubs, diamonds, hearts, and spades. It seems reasonable to just store these as strings.
We know there are 13 ranks: 2-10, jack, queen, king, ace. Integers would work for 2-10, but that doesn't work for the face cards.
We can just use strings as well to hold the rank value.
*/
type Card struct {
	Suit string
	Rank string
}

/*
Deck objects
------------
We know a deck of cards is just a collection of 52 cards containing one of each suit and rank.
With that in mind, we don't really need a Deck object; we can just treat a []Card as a deck. 
That seems to make things pretty simple.
*/

// NewDeck will create a standard, 52-card deck.  
func NewDeck() []*Card {
	suits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	ranks := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	cards := []*Card{}
	for _, suit := range suits {
		for _, rank := range ranks {
			card := &Card{suit, rank}
			cards = append(cards, card)
		}
	}

	return cards
}

/*
Shuffling
---------
We know users will want to shuffle their deck at some point, so we just add a helper function here
*/
  
func NewShuffledDeck() []*Card {
	deck := NewDeck()
	for i := 0; i < 52; i++ {
		j := rand.Intn(52)
		tmp := deck[i]
		deck[i] = deck[j]
		deck[j] = tmp
	}

	return deck
}

/*
Hands
-----
Hands are also just a collection of cards, so just like decks, hands can be expressed as just []*Card
*/
