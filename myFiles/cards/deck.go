package main

import (
	"fmt"
	"strings"
)

// Create a new type of desk
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// receiver function
// any variable of type "deck" now gets access to the print method
// d deck
// d - the actual copy of the deck we are working with is avaiable in the function as the variable d
// deck - variable of type deck
// refer to receiver as one or 2 letter- so deck- d
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// slices are zero-index
// go has helper built in
// [startIndexIncluding: upToNotIncluding]

// return 2 values of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
