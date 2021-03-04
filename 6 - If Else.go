package main

import "fmt"

func main6() {
	if 4+2 == 6 {
		fmt.Println("4 + 2 is indeed 6.")
	}

	// Note that the else has to be on the closing brace of the if!
	if 5/2 == 100 {
		fmt.Println("That's not right.")
	} else {
		fmt.Println("5 / 2 is not 100.")
	}
}
