package gollama

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type llamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type llamaResponse struct {
	Model      string `json:"model"`
	Response   string `json:"response"`
	DoneReason string `json:"done_reason"`
	Context    []int  `json:"context"`
	Done       bool   `json:"done"`
}

type OllamaClient struct {
	url   string
	model string
}

func New(url, model string) *OllamaClient {
	return &OllamaClient{
		url:   url,
		model: model,
	}
}

func (oc *OllamaClient) Generate(prompt string) (string, error) {
	client := http.DefaultClient
	query := &llamaRequest{
		Model:  oc.model,
		Prompt: prompt,
		Stream: false,
	}
	body, err := json.Marshal(query)
	if err != nil {
		return "", errors.New("Unable to generate request")
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/generate", oc.url), bytes.NewReader(body))
	if err != nil {
		return "", errors.New("Unable to generate request")
	}

	res, err := client.Do(req)
	if err != nil {
		return "", errors.New("Error sending request")
	}

	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", errors.New("Unable to read response")
	}

	response := &llamaResponse{}
	err = json.Unmarshal(bytes, response)
	if err != nil {
		return "", errors.New("Unable to parse ollama response")
	}
	return response.Response, nil
}
