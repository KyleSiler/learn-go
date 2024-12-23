package main

import (
	"context"
	"fmt"
	"net/http"
)

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
)

type logKey int

const (
	_ logKey = iota
	key
)

func main() {
	server := http.Server{
		Handler: Middleware(http.HandlerFunc(message)),
		Addr:    ":8080",
	}

	fmt.Println("Listening on 8080")

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}

func message(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	Log(ctx, Debug, "This is a debug statement")
	Log(ctx, Info, "This is an info statement")

	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("Ok"))
	if err != nil {
		fmt.Println(err)
	}
}

func ContextWithLevel(ctx context.Context, level Level) context.Context {
	return context.WithValue(ctx, key, level)
}

func LevelFromContext(ctx context.Context) (Level, bool) {
	level, ok := ctx.Value(key).(Level)
	return level, ok
}

func Log(ctx context.Context, level Level, message string) {
	fmt.Println("inside logger")
	var inLevel Level
	inLevel, ok := LevelFromContext(ctx)

	fmt.Printf("%s\n", inLevel)

	if !ok {
		return
	}

	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}
	if level == Info && (inLevel == Debug || inLevel == Info) {
		fmt.Println(message)
	}
}

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		// logLevel := req.Header["log_level"]
		logLevel := req.URL.Query().Get("log_level")

		ctx = ContextWithLevel(ctx, Level(logLevel))

		req = req.WithContext(ctx)

		handler.ServeHTTP(rw, req)
	})
}
