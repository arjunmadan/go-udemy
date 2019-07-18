package main

func main() {
	// cards := newDeck()

	// fmt.Println(cards.toString())
	cards := newDeckFromFile("blah.txt")
	cards.shuffle()
	//cards.saveToFile("blah.txt")
	cards.print()
	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	// cards.print()
}
