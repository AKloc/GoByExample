package main

import "fmt"

// Variadic functions let you pass as many parameters into a function
// as you want.

func sum(nums ...int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}

func main14() {
	ans := sum(1, 2, 3, 4, 5)

	if ans == 15 {
		fmt.Println("The answer was 15, you win!")
	}

	// YOU CAN PASS IN A SLICE, TOO.
	// Just use "..." after the slice.
	arr := []int{1, 2, 3, 4, 5}
	ans = sum(arr[1:2]...) // getting cute here and passing a slice of a slice.

	if ans == 15 {
		fmt.Println("The answer was 15, you win again!")
	} else {
		fmt.Println("Nope. Got back", ans)
	}
}
