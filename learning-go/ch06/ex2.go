package main

import "fmt"

func main() {
	slicey := []string{"one", "two"}

	UpdateSlice(slicey, "pizza")
	fmt.Println(slicey)

	GrowSlice(slicey, "burger")
	fmt.Println(slicey)
}

func UpdateSlice(big []string, little string) {
	big[len(big)-1] = little
	fmt.Println(big)
}

func GrowSlice(big []string, little string) {
	big = append(big, little)
	fmt.Println(big)
}
