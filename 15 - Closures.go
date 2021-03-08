package main

import "fmt"

// Here's a closure.
// This function returns another function which itself returns an int.
// What makes this a closure vs. an anonymous function? The i value persists
// and is unique to each instantiation of that function.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main15() {
	// "Give me a new function of intSeq"
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println("Kicking off the second batch")
	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println("Back to the other...")
	fmt.Println(nextInt())
}
