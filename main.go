package main

import "fmt"

func main(){
	// var card string = "Ace of Spades"
	card := newCard() //"Ace of Spades"  // := for initialization only
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}