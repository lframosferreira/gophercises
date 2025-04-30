package deck

import (
	"math/rand"
	"sort"
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

func ShuffleCards() func(*Deck) {
	return func(deck *Deck) {
		rand.Shuffle(len(deck.cards), func(i, j int) {
			deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
		})
	}
}

func AddJokers(amount int) func(*Deck) {
	return func(deck *Deck) {

		jokers := make([]Card, amount)
		for i := range amount {
			jokers[i] = Card{number: -1, suit: "joker"}
		}
		deck.cards = append(deck.cards, jokers...)
	}
}

func Sort(sort_func func(i, j int) bool) func(*Deck) {
	return func(deck *Deck) {
		sort.Slice(deck.cards, func(i, j int) bool {
			return deck.cards[i].number < deck.cards[j].number
		})
	}
}
