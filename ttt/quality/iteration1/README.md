# Iteration 1

# Before 
- what are my requirements?

# Review

## Does it meet all of my requirements?
Yes, I have a way to work with cards, decks, and hands. 
I have a helper function to create a new deck and to shuffle a deck of cards.

# What would consuming this code look like?
Let's use a simple blackjack scenario to test this out:

```go
func main() {
    // first, shuffle the deck and deal cards to each player
    deck := cards.NewSuffledDeck()
    for _, player := range players {
        firstCard := dealCard(deck)
        secondCard := dealCard(deck)
        player.Hand = []cards.Card{firstCard, secondCard}
    }

    // then, let each player play
    // we know each player may be dealt more cards during their turn, 
    // so the player.Play function would need to pass in the deck.  
    for _, player := range players {
        player.Play(deck)    
        player.Score = calculateScore(player.Hand)
    }

    // last, show who won
    winner := findWinner(players)
    fmt.Println("Winner is: %s", player.Name)
}

func dealCard(deck cards.Deck) *cards.Card {
    card = deck[0]
    deck = deck[1:]
    return card
}

func calculateScore(hand []*card.Card) int {
    // to keep things simple, we just say aces are 11
    values := map[string]int{
        "Two": 2,
        "Three": 3,
        ...
        "Jack": 10,
        "Queen": 10,
        "King": 10,
        "Ace": 11,  
    }

    var score int
    for _, card := range hand {
        score += values[card.Rank]
    }

    if score > 21 {
        return 0
    }

    return score
}
```

This exposes a couple problems with our implementation:
* Dealing from the deck is tedious
* Mapping a card's rank to a value is tedious   
* Treating decks and hands as []Card seems to work well

## Do my requirements need updating?
Yes!  
 
## What do I like? What do I dislike? What can I improve?  
Likes: treating decks and hands as []Card
Dislikes: tediousness around dealing cards and setting card values   
Improvements: How can I assign custom values to ranks? Pass in map[Rank]int? 

I typically write these things down and address them in the next iteration.
