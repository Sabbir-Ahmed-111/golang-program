package main

import "fmt"

func main() {

	//anoymous finction - je function name nai
	// Immediately Invoke Function Expression
	//IIFE

	func(a, b int) { //anonymous func je function ar name nai
		c := a + b
		fmt.Println(c)

	}(4, 5)
}

func init() {
	fmt.Println("I ' ll first executed")
}
