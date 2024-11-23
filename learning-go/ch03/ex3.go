package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	me := Employee{"Kyle", "Siler", 1}
	michelle := Employee{firstName: "Michelle", lastName: "Siler", id: 2}
	var john Employee
	john.firstName = "John"
	john.lastName = "Dorian"
	john.id = 3

	fmt.Println(me)
	fmt.Println(michelle)

	fmt.Println(john)
}
