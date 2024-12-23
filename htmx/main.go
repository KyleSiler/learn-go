package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		homePage, err := os.ReadFile("templates/home.html")
		if err != nil {
			log.Println("Couldn't read file")
		}
		resp.Write([]byte(homePage))
	})

	// TODO: Turn home.html into a template with information that gets passed in
	// Maybe make a basic htmx example where I pass information from a form and it gets saved. Like TODO app

	log.Print("Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
