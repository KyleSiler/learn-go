package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	e1, _ := fileLen("ex1.go")
	fmt.Println("ex1.go", e1)

	_, err := fileLen("wrong.go")
	fmt.Println("wrong", err)
}

func fileLen(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	data := make([]byte, 0, 2048)

	total := 0
	for {
		count, err := file.Read(data)

		total += count
		if err == io.EOF {
			return total, nil
		} else {
			return 0, err
		}
	}
}
