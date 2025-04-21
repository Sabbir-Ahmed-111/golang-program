package main

import "fmt"

/*func add(num1 int, num2 int) {
	sum := num1 + num2

	fmt.Println(sum)
}

func main() {
	a := 10
	b := 20

	add(a, b)

}
*/
/*
func add(num1 int, num2 int) int {
	sum := num1 + num2

	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2

	return sum, mul
}

func main() {
	a := 10
	b := 20

	p, q := getNumbers(a, b)

	fmt.Println(p)
	fmt.Println(q)
}
*/

func printWelcomeMessage() {
	fmt.Println("Welcome to the Application")
}
func getUserName() string {
	var name string
	fmt.Println("Enter Your Name -")
	fmt.Scanln(&name)
	return name
}

func getTwoNumbers() (int, int) {
	var num1 int
	var num2 int
	fmt.Println("Enter first number-")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number-")
	fmt.Scanln(&num2)
	return num1, num2
}

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func display(name string, sum int) {
	fmt.Println("Hello ", name)
	fmt.Println("summation ", sum)
}

func printGoodbyeMessage() {
	fmt.Println("Thank you for using Application ")
	fmt.Println("Goodbye")
}

func main() {
	printWelcomeMessage()
	name := getUserName()
	num1, num2 := getTwoNumbers()
	sum := add(num1, num2)
	display(name, sum)
	printGoodbyeMessage()

}
