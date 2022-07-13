Array and Slice are similar, except a slice is sized dynamically.  Both can only have a single data type

card := []string{}

func (d deck) print() {       (d deck) is an instance of a type set to receive the method
	for i, card := range d {
		fmt.Println(i, card)
		}
}