package main

import (
	"testing"
	"os"
)
func TestNewDeck(t *testing.T) {
	d := newDeck()
	// check if deck contains 16 cards
	if len(d) != 16 {
		t.Errorf("Expected deck length 16, but got %v", len(d))
	}
	// check if first item in deck is the expected value
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
	// check if last item in deck is the expected value
	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// remove any file named _decktesting pre-test
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	// check length of deck loaded from file
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}
	// remove any file named _decktesting post-test
	os.Remove("_decktesting")
}