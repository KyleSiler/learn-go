package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println(err.Error())
	}

	body := ""
	chunk := make([]byte, 10)
	for {
		_, err := resp.Body.Read(chunk)
		if err == io.EOF {
			break
		}

		body += string(chunk)
	}

	fmt.Println(body)
}
