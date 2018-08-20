package main

import "fmt"

// create a new type of 'deck'
// wich is a slice of strings
type deck []string

// Creates a new Deck, mixing to slices
// one with values and the other with suits
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit + " of " + value)
		}
	}
	return cards
}

// prints the deck on console
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}