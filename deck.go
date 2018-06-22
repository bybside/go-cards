package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
)

// create a new type of 'deck' which is a slice of strings
type deck []string

// roughly equivalent to a constructor
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, value := range cardValues {
		for _, suit := range cardSuits {
			cards = append(cards, value + " of "+ suit)
		}
	}

	return cards
}

// function with 'receiver' deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// instantiate a new hand with cards drawn from deck
func deal(d deck, handSize int) (deck, deck) {
	hand := d[:handSize]
	remainingDeck := d[handSize:]
	return hand, remainingDeck
}

func (d deck) toString() string {
	// takes a string slice and joins it into a single, comma delimited string
	return strings.Join([]string(d),",")
}

func (d deck) saveToFile(fileName string) error {
	// ioutil.WriteFile() writes a file to disk and takes
	// a file name, a byte slice of data, and permissions
	// for the file being written as arguments;
	// if an error occurs, returns it
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		// log the error and entirely quit the program
		fmt.Println("error:", err)
		os.Exit(1) // any value other than 0 indicates something went wrong while running the program
	}
	// cast byte slice as string and split (comma delimited) into string slice
	s := strings.Split(string(bs), ",")
	// return string slice cast as deck
	return deck(s)
}
