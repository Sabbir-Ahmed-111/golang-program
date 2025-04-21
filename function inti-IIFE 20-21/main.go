package main

import "fmt"

func main() {
	fmt.Println("I'am init function")
}

func init() {
	fmt.Println(" I am first function")
}

/*
var a = 10

func main(){
fmt.Println(a)

}

func init(){
fmt.Println(a)
a = 20
}

*/
