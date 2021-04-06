package main

import (
	"fmt"
	"time"
)

/*
Go’s select lets you wait on multiple channel operations.
Combining goroutines and channels with select is a powerful feature of Go.
*/

// Looks like a switch, but isn't! Basically if you put a channel receive in
//   it will evaluate everything and wait for the channels to receive before
//	 continuing.

// SELECT: Only used with channels, can choose multiple valid options as they
//   roll in.
// SWITCH: Only used with concrete types. Always go in sequence.

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
