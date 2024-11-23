package main

import "fmt"

func main() {
	var b byte = 255
	var littleI int32 = 2147483647
	var bigI int64 = 9223372036854775807

	b += 1
	littleI += 1
	bigI += 1

	fmt.Println(b)

	fmt.Println(littleI)
	fmt.Println(bigI)
}
