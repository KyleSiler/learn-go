package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2)*time.Second)

	defer cancel()

	sum, iterations, err := findNumber(ctx)

	var reason string
	if err != nil {
		reason = "Timeout"
	} else {
		reason = "Found 1234"
	}
	fmt.Printf("sum %d - iter %d - reas: %s", sum, iterations, reason)
}

func findNumber(ctx context.Context) (int, int, error) {
	sum, iterations := 0, 0
	for {
		n := rand.Intn(100_000_000)

		sum += n
		iterations++

		if n == 1234 {
			return sum, iterations, nil
		}

		select {
		case <-ctx.Done():
			return sum, iterations, ctx.Err()
		default:
		}
	}
}
