package main

import "fmt"

const a = 10

var p = 100

func outer() func() {
	money := 100
	age := 30
	fmt.Println("Age :", age)

	show := func() {
		money = money + a + p
		fmt.Println(money)
	}

	return show

}

func call() {
	inr1 := outer() //show
	inr1()
	inr1()

	inr2 := outer()
	inr2()
	inr2()

}
func main() {
	call()
}

func init() {
	fmt.Println("---Bank---")
}
