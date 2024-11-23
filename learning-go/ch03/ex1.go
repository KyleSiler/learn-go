package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "AnotherGreeting", "Konichiwa", "Como estas"}

	firstTwo := greetings[:2]

	middle := greetings[1:4]

	fourthFifth := greetings[3:]

	fmt.Println(firstTwo)

	fmt.Println(middle)

	fmt.Println(fourthFifth)
}
