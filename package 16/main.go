package main

import (
	"fmt"

	"example.cm/mathlib"
)

var (
	a = 10
	b = 20
)

func main() {
	fmt.Println("showing custom package")
	mathlib.Add(4, 7)
	mathlib.Money(50, 25)
}
