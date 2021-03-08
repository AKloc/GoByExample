package main

import "fmt"

// We can add methods onto structs, Python-style.
// Just add the type prior to the method name - this is called the "receiver"
type rect struct {
	width, height int
}

func newRect(width, height int) *rect {
	r := rect{width: width, height: height}
	return &r
}

// Here's a method. Recommended to generally use pointer receivers.
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main19() {
	r := newRect(5, 10)
	fmt.Println(r)
}
