package main

import "fmt"

func main() {
	// s2 := []int{1, 2, 5} //slice literal
	// fmt.Println("slice:", s2, "len:", len(s2), "cap:", cap(s2))

	// s3 := make([]int, 3) //make func with len

	// s3[0] = 5

	// fmt.Println(s3)

	// fmt.Println(len(s3))
	// fmt.Println(cap(s3))

	// s4 := make([]int, 3, 5) //make func with len and capacity

	// s4[0] = 5
	// s4[2] = 10

	// fmt.Println(s4)

	// fmt.Println(len(s4))
	// fmt.Println(cap(s4))

	var s5 []int       //empty slice or nill slice
	s5 = append(s5, 3) //append dia value add kore
	fmt.Println(s5)

}
