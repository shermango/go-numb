package main

import (
	"os"
	"testing"
)

// test newDeck
func TestNewDeck(t *testing.T) {
	d := newDeck()

	// should init with 52 cards
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, instead received actual %v", len(d))
	}

	// first card should be ace of spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card in new deck to be Ace of Spades, instead received actual %v", d[0])
	}

	// last card should be king of clubs
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card in new deck to be King Of Clubs, instead received %v", d[len(d)-1])
	}
}

// test saveToFile and loadDeckFromFile
func TestSavingAndLoadingDecks(t *testing.T) {
	// init testing environment from scratch
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	// should save the deck
	_, err := os.Open("_decktesting")
	if err != nil {
		t.Errorf("Expected deck to be saved but instead received error %v ", err)
		os.Exit(1)
	}

	// should load the saved deck
	loadedDeck := loadDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in loaded deck, instead received actual %v ", len(loadedDeck))
	}

	// reset testing enviroment
	os.Remove("_decktesting")
}

// test Deal
func TestDeal(t *testing.T) {
	d := newDeck()
	hand, remainingDeck := deal(d, 5)

	if len(hand) != 5 {
		t.Errorf("Expected dealt hand to have size 5, instead received actual %v ", len(hand))
	}

	if len(remainingDeck) != 47 {
		t.Errorf("Expected dealt hand to have size 47, instead received actual %v ", len(remainingDeck))
	}
}

// test shuffle
func TestShuffle(t *testing.T) {
	d := newDeck()
	firstCard := d[0]
	d.shuffle()
	firstCardAfterShuffle := d[0]

	// this is probably really bad logic to determine a shuffle
	if firstCard == firstCardAfterShuffle {
		t.Fail()
	}

	// not sure if I should be saving a reference or not, but it's like 1am and I really don't care about this too much
}
