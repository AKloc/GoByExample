package main

import "fmt"

func main7() {
	i := 2
	fmt.Println("i is set to", i)

	switch i {
	case 1:
		fmt.Println("One.")
	case 2:
		fmt.Println("Two.")
	case 3:
		{
			fmt.Println("Three.")
			fmt.Println("Did the braces work here?")
		}
	default:
		fmt.Println("No idea what that number is.")
	}
}
