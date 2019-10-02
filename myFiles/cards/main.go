package main

import "fmt"

func main() {
	// assigns a variable to value card
	// var creates a new variable
	// card is name of variable
	// string- only a string will ever be assing to this value
	// ace of spaces is string that gets applied to card

	// go is statically typed- we care what we pass in
	// javascript is dynamically typed

	// 2 lines below are exactly the same
	// go will infer

	// var card string = "Ace of Spades"
	// only need : when first initialize
	// card := "Ace of Spades"

	// next time can do card = "Queen", etc no :

	// creating a slice
	cards := newDeck()

	// for index, card
	// index of element in array
	// card- current card
	// range cards- loop through cards
	// for loops declared everytime that's why :=

	fmt.Println(cards.toString())
}

// basic go types
// bool, string, int, float64

// Variables can be initialized outside of a function, \
// but cannot be assigned a variable.

// tells go compiler to return type string
func newCard() string {
	return "Five of Diamonds"
}

// Array - fixed length list of things
// Slice - an array that can grow or shrink
// every element in a slice must be of the same type
