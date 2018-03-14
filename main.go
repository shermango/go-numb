package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	fmt.Println("You probably wanna be checking out deck.go instead")
}
