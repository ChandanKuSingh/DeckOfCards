package main

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
	//fmt.Println(cards.toString())
	cards.saveToFile("myCards")
	cards = newDeckFromFile("myCards")
	cards.print()
	cards.shuffle()
	cards.print()
	cards.print()
}
