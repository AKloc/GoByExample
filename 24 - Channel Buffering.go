package main

import "fmt"

/*
By default channels are unbuffered, meaning that they will only accept sends (chan <-)
if there is a corresponding receive (<- chan) ready to receive the sent value.
Buffered channels accept a limited number of values without a corresponding receiver for those values.
*/

func main() {

	// Note the "2" parameter - that means that we can buffer two values at a time.
	messages := make(chan string, 2)
	messages <- "Here's the first"
	messages <- "Here's the second"
	// Below will deadlock because our buffer isn't big enough...
	//messages <- "How about a third? I'm guessing it'll show the second and third messages."

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	// ... even if we try receiving a third time.
	//fmt.Println(<-messages)
}
