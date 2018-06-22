package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("deck.txt")
	cards := newDeckFromFile("deck.txt")
	cards.print()
}
