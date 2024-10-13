package main

import (
	"log"
	"net/http"
)

func handleTest(resp http.ResponseWriter, request *http.Request) {
	resp.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/test", handleTest)
	port := ":8080"
	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
