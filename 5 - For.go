package main

import (
	"fmt"
)

func main5() {
	i := 1

	// Straightforward loop.
	for i <= 3 {
		fmt.Println("On loop", i)
		i++
	}

	// More traditional for loop.
	for j := 7; j <= 10; j++ {
		fmt.Println("More formal loop, iteration", j)
	}

	// Infinite loop / break
	for {
		fmt.Println("Infinite loop until I break!")
		break
	}
}
