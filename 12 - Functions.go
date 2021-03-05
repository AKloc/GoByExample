package main

import "fmt"

func main12() {
	x, y := 1, 2
	res := plus(x, y)
	fmt.Printf("%d + %d = %d\n", x, y, res)
}

func plus(x, y int) int {
	return x + y
}
