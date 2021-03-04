package main

import (
	"fmt"
)

const s = "constant"

func main4() {
	fmt.Println(s)

	const n = 50000000

	fmt.Println(s, n)
}
