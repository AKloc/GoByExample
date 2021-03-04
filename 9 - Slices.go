package main

import "fmt"

func main9() {
	// Slices are a huge part of Go. Used more often than arrays (like Python)
	// Slices are NOT just parts of arrays, they're more like lists. Can
	// change their size dynamically.

	// To create, we have to use "make".

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("After setting, emp =", s)

	// NOW - let's add on with "append"
	s = append(s, "d")
	s = append(s, "e")
	fmt.Println("After appending two values, emp = ", s, "and has length", len(s))

	// copy works, a little weird. takes two params, first is destination, second is source.
	c := make([]string, len(s))
	copy(c, s)

	fmt.Println("Here's c.", c)

	// Slice operator works same as Python.
	// First index is where you start, second is where you end. Open / closed.
	l := c[3:]
	fmt.Println("Last few...", l)

	m := c[2:4]
	fmt.Println("Middle two...", m)
}
