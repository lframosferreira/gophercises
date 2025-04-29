package deck

import (
	"math/rand"
)

type Card struct {
	number int
	suit   string
}

func NewCard(number int, suit string) *Card {
	card := Card{number: number, suit: suit}
	return &card
}

type Deck struct {
	cards []Card
}

func New(options ...func(*Deck)) Deck {
	deck := Deck{}
	for _, suit := range []string{"spades", "diamonds", "clubs", "hearts"} {
		for number := range 14 {
			deck.cards = append(deck.cards, Card{number: number, suit: suit})
		}
	}

	for _, opt := range options {
		opt(&deck)
	}
	return deck
}

func ShuffleCards(should_shuffle bool) func(*Deck) {
	return func(deck *Deck) {
		if should_shuffle {
			rand.Shuffle(len(deck.cards), func(i, j int) {
				deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
			})
		}
	}
}
