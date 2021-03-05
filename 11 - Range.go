package main

import "fmt"

func main11() {
	// "range" iterates over elements.
	// gives two values back - key and value.

	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		fmt.Println("Looking at", num)
		sum += num
	}

	fmt.Println("sum:", sum)

	// Does it work with maps? SURE DOES.
	m := map[string]int{"one": 1, "two": 2, "three": 3}

	for k, v := range m {
		// printf verbs: %v uses default style. %s is string. %q wraps in quotes.
		// %d is digits, %t is true / false.
		fmt.Printf("%s -> %d\n", k, v)
		fmt.Printf("Key:%v, Value:%v\n", k, v)
	}

}
