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

	fmt.Println(nums)
}
