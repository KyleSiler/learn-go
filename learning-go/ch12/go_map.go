package main

import (
	"fmt"
	"math"
	"sync"
)

func buildCacheMap() map[int]float64 {
	myMap := make(map[int]float64)

	for i := 0; i < 100_000; i++ {
		myMap[i] = math.Sqrt(float64(i))
	}

	return myMap
}

var squareRootMapCache = sync.OnceValue(buildCacheMap)

func main() {
	for i := 0; i < 100_000; i += 1_000 {
		fmt.Println(squareRootMapCache()[i])
	}
}
