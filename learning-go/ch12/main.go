package main

import "fmt"

func main() {
	ch := make(chan int, 20)
	go func() {
		for i := range 10 {
			ch <- i
		}
	}()
	go func() {
		for i := 20; i < 30; i++ {
			ch <- i
		}
	}()

	for i := 0; i < 20; i++ {
		fmt.Println(<-ch)
	}
	close(ch)
}
