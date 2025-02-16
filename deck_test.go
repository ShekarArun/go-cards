package main

import (
	"os"
	"testing"
)

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

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Delete any existing test files
	os.Remove("_decktesting")

	// Seed test data
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	// Run test
	deckLength := 52
	if len(loadedDeck) != deckLength {
		t.Errorf("Expected deck length of %v, but got %v", deckLength, len(loadedDeck))
	}

	// Cleanup any files created
	os.Remove("_decktesting")
}
