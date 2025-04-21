package main

import "fmt"

type user struct {
	Name string //member variable or property
	Age  int
}

func UserDetails(usr user) { //regulation function
	fmt.Println("Name :", usr.Name)
	fmt.Println("Age :", usr.Age)
}

func (usr user) printDetails() { // receiver function
	fmt.Println("Name :", usr.Name)
	fmt.Println("Age :", usr.Age)
}

func main() {

	var user1 user

	user1 = user{ // Instance or Object
		Name: "Sabbir",
		Age:  25,
	}

	user1.printDetails()

	//UserDetails(user1)

	user2 := user{
		Name: "Roki",
		Age:  26,
	}

	//UserDetails(user2)

	user2.printDetails()

}
