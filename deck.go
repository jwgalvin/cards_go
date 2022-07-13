package main
// create a new 'type' of called deck. it is a slice of strings. Similar to a class
import "fmt"

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
		}
}