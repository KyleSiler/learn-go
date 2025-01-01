package main

import (
	"log"
	"ollama-client-go/gollama"
)

func main() {
	client := gollama.New("http://oasis.local:11434", "llama3.2")
	output, err := client.Generate("What color is the sky?")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(output)
}
