package main

import "fmt"

// Passing by value, so ival will NOT be modified.
func zeroval(ival int) {
	ival = 0
}

// Passing by reference, so iptr WILL be modified.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main17() {
	i := 1
	fmt.Println("Initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	// This will display the actual value.
	fmt.Println("zeroptr:", i)

	// This will display the memory address.
	fmt.Println("pointer:", &i)
}
