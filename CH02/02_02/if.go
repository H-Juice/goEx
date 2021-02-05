package main

import (
	"fmt"
)

func main() {

	x := 10

	if x > 5 {
		fmt.Println("x is big")
	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more tha half of b")
	}

}
