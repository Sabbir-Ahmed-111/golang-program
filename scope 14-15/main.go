package main

import "fmt"

var a = 20
var b = 20

func add(x int, y int) {

	z := x + y
	fmt.Println(z)
}

func main() {
	p := 40
	q := 50

	add(a, p)

	add(a, b)

	add(p, q)
}
