package main

import "fmt"

func main() {
	// TODO: Implement two channels here
	//
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 100; i < 110; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	openChannels := 2
	for openChannels != 0 {
		select {
		case i, ok := <-ch1:
			if !ok {
				openChannels--
				ch1 = nil
			} else {
				fmt.Println(i)
			}
		case j, ok := <-ch2:
			if !ok {
				openChannels--
				ch2 = nil
			} else {
				fmt.Println(j)
			}
		}
	}
}
