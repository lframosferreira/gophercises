package main

import (
	"deck_of_cards/deck"
	"fmt"
)

func main() {

	deck := deck.New(deck.ShuffleCards(true))
	fmt.Print(deck)
}
