package main

import (
	"fmt"
	"math"
)

// Interfaces are implemented implicitly - kind of neat.
// So anything that implements area() and perim() will have implemented
//   the geometry interface below.
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (r *rect) perim() float64 {
	return (2 * r.width) + (2 * r.height)
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Here's the magic!
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perim())
}

func main20() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(&r)
	measure(&c)
}