package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	persons := make([]Person, 0)
	for range 10_000_000 {
		persons = append(persons, Person{FirstName: "Kyle", LastName: "Siler", Age: 5})
	}
}
