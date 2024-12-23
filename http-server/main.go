package main

import (
	"log"
	"net/http"
)

func handleTest(resp http.ResponseWriter, request *http.Request) {
	resp.Write([]byte("Hello World"))
}

type MyHandler struct{}

func (h MyHandler) ServeHTTP(resp http.ResponseWriter, request *http.Request) {
	resp.Write([]byte("Another one"))
}

func main() {
	mux := http.NewServeMux()

	test := http.NewServeMux()

	test.Handle("/test", MyHandler{})

	mux.Handle("/anotherone/", http.StripPrefix("/anotherone", test))

	mux.HandleFunc("/test", handleTest)
	port := ":8080"
	log.Printf("Starting server on port %s", port)
	// log.Fatal(http.ListenAndServe(port, mux))

	server := http.Server{
		Addr:    port,
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
