package main

import "fmt"

func main3() {
	var a = "initial"
	fmt.Println(a)

	// := declares and initializes in one step.
	b := "second"
	fmt.Println(b)

	var c, d int = 2, 3
	fmt.Println(c, d)
}
