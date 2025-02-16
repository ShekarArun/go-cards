package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	deckLength := 52
	if len(d) != deckLength {
		t.Errorf("Expected deck length of %v, but got %v", deckLength, len(d))
	}

	firstCard := "Ace of Spades"
	if d[0] != firstCard {
		t.Errorf("Expected first card to be %v, but got %v", firstCard, d[0])
	}

	lastCard := "King of Clubs"
	if d[len(d)-1] != lastCard {
		t.Errorf("Expected last card to be %v, but got %v", lastCard, d[len(d)-1])
	}
}
