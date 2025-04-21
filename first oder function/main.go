package main

import "fmt"

func processOperation(a int, b int, op func(p int, q int)) {
	op(a, b)

}

//

func dd(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	processOperation(3, 4, dd)
}
