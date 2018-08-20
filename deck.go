package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)


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
			cards = append(cards, value + " of " + suit)
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

// returns a hand of cards, and the rest of cards on the deck 
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// converts a deck to a single string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// saves the deck on a file with the specified name
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 7)
}

// reads a file and gets the data and converts it to a deck
func newDeckFromFile(filename string) deck{
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs),","))
}