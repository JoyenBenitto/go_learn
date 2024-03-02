package main

import "fmt"

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
