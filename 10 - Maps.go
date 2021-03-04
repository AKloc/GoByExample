package main

import "fmt"

func main10() {
	// Key value pairs.

	// Creates a map with a string key and an int value
	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2

	fmt.Println("Map's length is", len(m), "value:", m)

	// this misses, but returns a default.
	fmt.Println("Map[\"oned\"] =", m["oned"])

	// OR, we can specify the second return value, which returns whether or not
	// the key was found.
	v, prs := m["one"]
	if prs {
		fmt.Println("Found that value, and it is", v)
	}

	m["three"] = 3
	// get rid of values with "delete"
	delete(m, "three")

	// Here's how we can define a map and put values in it right away.
	// NOTE - don't have to use "make" here.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("Here's our second map.", n)
}
