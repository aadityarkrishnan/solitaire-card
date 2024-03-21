package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDecks()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the first card to be Ace of Spades, But received  %v", d[0])

	}

	if d[len(d)-1] != "Four of Calve" {
		t.Errorf("Expected the last card to be Four of Calve, But received  %v", d[len(d)-1])

	}
}

func TestSaveDecktoReadDeck(t *testing.T) {
	os.Remove("_testing_deck")
	deck := newDecks()
	deck.saveToFile("_testing_deck")

	loaded_deck := newDeckFromFile("_testing_deck")

	if len(loaded_deck) != 16 {
		t.Errorf("Expected 16 cards, but received only %v", len(loaded_deck))
	}

	os.Remove("_testing_deck")

}
