package main

// create a new 'type' of called deck. it is a slice of strings. Similar to a class
import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string


//no receiver as this creates the deck

func newDeck() deck {   
	cards := deck{}
	
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	
	for _, suit :=  range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
		}
}
  // name(var name 'type', var 'type) return type
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}