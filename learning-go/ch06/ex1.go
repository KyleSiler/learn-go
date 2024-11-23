package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	person := MakePerson("Kyle", "Siler", 36)
	pointer := MakePersonPointer("Michelle", "Siler", 34)

	fmt.Println(person)
	fmt.Println(pointer)
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{FirstName: firstName, LastName: lastName, Age: age}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{FirstName: firstName, LastName: lastName, Age: age}
}
