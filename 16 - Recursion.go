package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return fact(n-1) * n
}

func main16() {
	o := 5
	f := fact(o)
	fmt.Printf("Factorial of %d is %d", o, f)
}
