package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Double[T Number](a T) T {
	return a * 2
}

// Write a generic function that doubles the value of any integer or float passed in
func main() {
	fmt.Println(Double[int](int(5)))
	fmt.Println(Double(float32(6)))
}
