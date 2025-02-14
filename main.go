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

	fmt.Println(hand.saveToFile("hand.txt"))

	deck2 := newDeckFromFile("hand.txt")
	fmt.Println("Deck2:")
	deck2.print()

	fmt.Println("Intentionally loading a non-existent file:")
	deck3 := newDeckFromFile("non-existent.txt")
	deck3.print()
}
