package main

import (
	"errors"
	"log"
)

func add(i int, j int) (int, error) {
	return i + j, nil
}

func sub(i int, j int) (int, error) {
	return i - j, nil
}

func mul(i int, j int) (int, error) {
	return i * j, nil
}

func div(i int, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	}
	return i / j, nil
}

var opMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	_, err := opMap["/"](5, 0)
	if err != nil {
		log.Fatal(err)
	}
}
