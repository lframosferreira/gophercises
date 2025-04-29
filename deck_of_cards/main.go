package main

import (
	"deck_of_cards/deck"
	"fmt"
)

func main() {

	deck := deck.New(deck.AddJokers(5), deck.ShuffleCards())
	fmt.Print(deck)
}
