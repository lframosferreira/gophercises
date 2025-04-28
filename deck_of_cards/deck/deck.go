package deck

type Card struct {
	number int
	suit   string
}

func new_card(number int, suit string) *Card {
	card := Card{number: number, suit: suit}
	return &card
}

func New() []Card {
	cards := []Card{}
	for _, suit := range []string{"spades", "diamonds", "clubs", "hearts"} {
		for number := range 14 {
			cards = append(cards, Card{number: number, suit: suit})
		}
	}

	return cards
}
