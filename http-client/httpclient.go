package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CatFact struct {
	Text string
	Type string
}

func main() {
	resp, err := http.Get("https://cat-fact.herokuapp.com/facts/random")
	if err != nil {
		fmt.Println(err.Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	// fact := new(CatFact)
	var fact CatFact
	json.Unmarshal(body, &fact)

	fmt.Println(fact.Text)
	fmt.Println("")
	fmt.Println(string(body))
}
