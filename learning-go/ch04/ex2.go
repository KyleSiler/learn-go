package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var nums []int
	for range 100 {
		nums = append(nums, rand.Intn(100))
	}

	for _, v := range nums {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six!", v)
		case v%2 == 0:
			fmt.Println("Two!", v)
		case v%3 == 0:
			fmt.Println("Three!", v)
		default:
			fmt.Println("Never mind", v)
		}
	}
}
