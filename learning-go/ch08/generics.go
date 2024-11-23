package main

import (
	"fmt"
	"strconv"
)

func Map[T1, T2 any](ts []T1, mapper func(T1) T2) []T2 {
	output := make([]T2, len(ts))

	for i, k := range ts {
		output[i] = mapper(k)
	}

	return output
}

func Reduce[T1, T2 any](input []T1, init T2, reducer func(T2, T1) T2) T2 {
	r := init
	for _, v := range input {
		r = reducer(r, v)
	}
	return r
}

func Filter[T any](input []T, filter func(T) bool) []T {
	var output []T

	for _, v := range input {
		if filter(v) {
			output = append(output, v)
		}
	}

	return output
}

func main() {
	myints := []int{1, 2, 3, 4, 5}

	output := Map(myints, func(t int) string {
		t += 1
		return strconv.Itoa(t)
	})

	fmt.Println(myints)
	fmt.Println(output)

	reducer := Reduce(myints, "Hi ints", func(init string, t int) string {
		return fmt.Sprint(init, " ", t)
	})

	fmt.Println(reducer)

	filter := Filter(myints, func(myint int) bool {
		return myint%2 == 0
	})

	fmt.Println(filter)
}
