package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*Create  a new type 'deck'
which is a slice of strings
*/

type deck []string

func newDeck() deck {
	cardSuits := []string{"clubs", "Heart", "Spade", "Diamonds"}
	cardValue := []string{"Ace", "1", "2", "3", "4", "5"}
	cards := deck{}

	for _, suit := range cardSuits {
		for _, val := range cardValue {
			cards = append(cards, val+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckfromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) //byteslice and error
	if err != nil {
		//option log the error and return call to a new deck
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), "s")
	return deck(s)
}
