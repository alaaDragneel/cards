package main

func main() {
	// var card string = "Ace of Spades"

	// By Using ":=" We tell The Go Compiler To Define A New Variable And Figure Out What It's Type and use it when Defining New Variable Only
	// card := newCard()
	// fmt.Println(card)

	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
