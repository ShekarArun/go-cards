package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards2 := append(cards, "Six of Spades") // Note that the original slice is not modified

	// fmt.Println(cards)
	// fmt.Println(cards2)

	// cards.print()

	deck1 := newDeck()
	deck1.print()

	hand, remainingDeck := deal(deck1, 5)
	fmt.Println("Hand:")
	hand.print()
	fmt.Println("Remaining deck:")
	remainingDeck.print()

	fmt.Println(hand.toString())
}
