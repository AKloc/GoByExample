package main

import (
	"fmt"
)

/*
Channels are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine and receive those values into another goroutine.

IMPORTANT: BOTH SENDS AND RECEIVES BLOCK UNTIL BOTH SENDER AND RECEIVER ARE READY.
*/

func main23() {

	// First we create a copmmunication channel that works with strings.
	messages := make(chan string)

	go func() {

		// If you uncomment this, you won't see the message at the end display for another 5 seconds.
		//time.Sleep(5000 * time.Millisecond)

		// Send the value "ping" into the channel ONCE THE RECEIVER IS READY
		messages <- "ping"

	}()

	// msg is just a regular string. Dump the contents of the channel into it.
	// This won't execute until the sending statement above is ready.
	msg := <-messages

	//msg := "What if there isn't a receiver? Will I see this? My guess is YES."
	fmt.Println(msg)
}
