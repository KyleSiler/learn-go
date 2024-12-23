package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	middleware := middlewareGenerator(100)
	server := http.Server{
		// TODO: Implement sleepy which is a func sleepy
		Handler: middleware(http.HandlerFunc(sleepy)),
		Addr:    ":8080",
	}
	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

func sleepy(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	message, err := doThing(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			rw.WriteHeader(http.StatusGatewayTimeout)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		rw.WriteHeader(http.StatusOK)
	}

	_, err = rw.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func doThing(ctx context.Context) (string, error) {
	wait := rand.Intn(200)

	select {
	case <-time.After(time.Duration(wait) * time.Millisecond):
		return "Done!", nil
	case <-ctx.Done():
		return "Too slow!", ctx.Err()
	}
}

func middlewareGenerator(milis int) func(http.Handler) http.Handler {
	return func(hand http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			ctx, cancel := context.WithTimeout(ctx, time.Duration(milis)*time.Millisecond)
			defer cancel()
			req = req.WithContext(ctx)
			hand.ServeHTTP(rw, req)
		})
	}
}
