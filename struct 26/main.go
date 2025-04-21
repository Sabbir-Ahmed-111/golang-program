package main

import "fmt"

type user struct {
	Name string //member variable or property
	Age  int
}

func main() {

	var user1 user

	user1 = user{ // Instance or Object
		Name: "Sabbir",
		Age:  25,
	}
	fmt.Println("Name :", user1.Name)
	fmt.Println("Age :", user1.Age)

	user2 := user{
		Name: "Roki",
		Age:  26,
	}

	fmt.Println("Name :", user2.Name)
	fmt.Println("Age: ", user2.Age)
}
