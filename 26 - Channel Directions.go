package main

import "fmt"

/*
When using channels as function parameters, you can specify if a channel is
meant to only send or receive values.
This specificity increases the type-safety of the program.
*/

// BASICALLY: chan<- means "a channel that can only send stuff"
//			  <- chan means "a channel that can only receive stuff"
// You don't HAVE to use these

// takes a channel for ONLY sending.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// takes a channel for ONLY receiving.
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main26() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
