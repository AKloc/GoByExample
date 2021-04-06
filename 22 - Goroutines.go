package main

import (
	"fmt"
	"time"
)

// Goroutines are really straightforward. Literally just add "go" before a function call.
// Can call anonymous functions the same way.
func f(from string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond * 250)
		fmt.Println(from, ":", i)
	}
}

func main22() {
	// Regular method call.
	f("direct call")

	// Calling that same method concurrently!
	go f("Goroutine call")

	// Calling an anonymous method concurrently. Note how you define
	// the method and then pass the parameters in right after it to run it.
	// This is probably very useful
	go func(msg string) {
		fmt.Println(msg)
	}("anonymous function going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
