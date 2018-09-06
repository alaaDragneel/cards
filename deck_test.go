package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// 16 Because We Have 4 Suites And 4 Values Go And See Deck.go newDeck()
	if len(d) != 16 {
		t.Errorf("Expected Deck Length Of 16, But Got %v", len(d))
	}

	if d[0] != "Ace Of Spades" {
		t.Errorf("Expected First Card Of Ace Of Spades, But Got %v", d[0])
	}

	if d[len(d)-1] != "Four Of Clubs" {
		t.Errorf("Expected Last Card Of Four Of Clubs, But Got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_decktesting"

	os.Remove(fileName)

	deck := newDeck()

	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected Deck Length Of 16, But Got %v", len(loadedDeck))
	}

	os.Remove(fileName)
}
