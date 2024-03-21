package main

import "fmt"

func main() {
	cards := newDecks()
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()

	cards1 := newDecks()
	fmt.Println(cards1.toString())

	cards2 := newDecks()
	cards2.saveToFile("playing_card")

	cards3 := newDeckFromFile("playing_card")
	cards3.print()

	cards4 := newDecks()
	cards4.shuffle()
	cards4.print()

}
