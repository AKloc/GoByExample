package main

import "fmt"

func main8() {
	// Declare an array, a, of size 5 ints.
	var a [5]int

	// Prints [0 0 0 0 0]
	fmt.Println("Here's an empty array.", a)

	// len(array) to get size of array, just like Python.
	fmt.Println("The size of our array is", len(a))

	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[len(a)-1] = 5
	fmt.Println("After modifying a few values:", a)

	fmt.Println("Second item is", a[1])

	for i := 0; i < len(a); i++ {
		fmt.Println("Item at index", i, "is:", a[i])
	}

	// Making a 2d array
	var twoD [5][10]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			twoD[i][j] = j * i
		}
	}

	fmt.Println("2D array:", twoD)
}
