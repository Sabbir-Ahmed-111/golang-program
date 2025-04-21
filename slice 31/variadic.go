package main

import "fmt"

//variadic function
//aita go ar building function
//akhane numbers slice
func print(numbers ...int) {
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}

func main() {
	print(1, 2, 3, 4, 5)

}
