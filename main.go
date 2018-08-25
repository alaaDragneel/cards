package main

func main() {
	// var card string = "Ace of Spades"

	// By Using ":=" We tell The Go Compiler To Define A New Variable And Figure Out What It's Type and use it when Defining New Variable Only
	// card := newCard()
	// fmt.Println(card)

	cards := deck{"Ace of Diamonds", newCard()}

	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
