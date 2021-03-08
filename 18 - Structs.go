package main

import "fmt"

// Structs. Classes, sans methods.
type person struct {
	name string
	age  int
}

// This is a common idiom in Go when creating a new instance.
// newWhatever returns a reference to the new Whatever.
func newPerson(name string) *person {
	p := person{name: name, age: 42}
	return &p
}

func main18() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})

	s := newPerson("Mikey")
	fmt.Println(s)

	fmt.Println("Mikey's age is", s.age)
}
