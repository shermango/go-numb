package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// OOP in Go?! call me the objective-g boi lul ¯\_(ツ)_/¯

/*
	@desc deck type is a string slice representing a deck of playing cards, it
	includes custom methods to perform traditional playing card operations,
	i.e. deal cards, shuffle cards, cheat the house, etc.
*/

type deck []string

// deck methods

/*
	@name newDeck
	@desc deck type initializer - creates a new deck
	@return {deck} returns a new instance of deck type
*/

func newDeck() deck {
	_newDeck := deck{}

	// define suits and values for the cards
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// loop through suits and values to create a unique card, i.e. Ace of Clubs
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := value + " of " + suit
			_newDeck = append(_newDeck, card)
		}
	}

	return _newDeck
}

/*
	@name print
	@desc prints out remaining cards in deck
	@receiver {deck} - the instance of dec to print
*/

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
	return
}

/*
	@name deal
	@desc deals a hand up to the given handSize
	@params {deck}
	@params {int} handSize the hand size to deal
	@return {deck} deal everything up to the hand size
	@return {deck} return the remaining deck
*/

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/*
	@name shuffle
	@desc basic shuffling of the deck, definitely not fisher-yates
	@receiver {deck} the deck instance
*/

func (d deck) shuffle() {
	/* I can't believe how convoluted this (generating a random) was, truly go
	programmers are truly called gophers based on the number of holes they dig
	through trying to navigate the official docs */

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// utility functions for deck type

/*
	@name toString
	@desc utility function to convert a deck type into a string type
	@receiver {deck} the instance of deck type to convert into string type
	@return {string} returns a string representation of the deck
*/

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

/*
	@name saveToFile
	@desc rudimentary method to save a deck instance to drive as a .txt file
	@receiver {deck} the instance of deck to save
	@return {error} returns error if error
*/

func (d deck) saveToFile(fileName string) error {
	// 0666 = read / write permissions
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

/*
	@name newDeckFromFile
	@desc rudimentary deck creation from a saved .txt file
	@return {deck} returns the deck
*/

func loadDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}
