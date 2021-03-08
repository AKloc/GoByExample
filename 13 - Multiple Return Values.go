package main

import "fmt"

func main13() {
	x, y := multipleVals()
	fmt.Printf("Got back two numbers. %d and %d\n", x, y)
	fmt.Println(singleVal())

	_, z := multipleVals()
	fmt.Println("Eh, I don't care about x, but y is", z)
}

// Returns two ints. Note that you have to wrap the return values in parens,
// unlike a single value.
func multipleVals() (int, int) {
	return 3, 5
}
func singleVal() int {
	return 2
}
