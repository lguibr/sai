package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const OpenAIAPIURL = "https://api.openai.com/v1/chat/completions"

type OpenAIRequest struct {
	Model    string                   `json:"model"`
	Messages []map[string]interface{} `json:"messages"`
}

func main() {
	history, err := getLastTenBashCommand()
	if err != nil {
		fmt.Printf("Error getting bash history: %v\n", err)
		return
	}

	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current path: %v\n", err)
		return
	}

	operatingSystem := runtime.GOOS

	var userMessage string
	if len(os.Args) > 1 {
		userMessage = strings.Join(os.Args[1:], " ")
	} else {
		userMessage = "Given my last 10 commands and the user operational system and the current path what is the most likely command to be sent:"
	}

	systemMessage := fmt.Sprintf("You are a bash completion tool and will receive a question and should answer it with a raw text with a valid command knowing that the user is in the path %s, is using the OS: %s, and the last 10 commands sent are:\n%s DON'T FORMAT OR ADD CODE BLOCK JUST RAW TEXT SINGLE LINE!", currentPath, operatingSystem, history)

	requestPayload := OpenAIRequest{
		Model: "gpt-4",
		Messages: []map[string]interface{}{
			{
				"role":    "system",
				"content": systemMessage,
			},
			{
				"role":    "user",
				"content": userMessage,
			},
		},
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	response, err := sendOpenAIRequest(requestPayload, apiKey)
	if err != nil {
		fmt.Printf("Error sending request to OpenAI: %v\n", err)
		return
	}

	var responseMap map[string]interface{}
	if err := json.Unmarshal([]byte(response), &responseMap); err != nil {
		fmt.Printf("\n\nError parsing response from OpenAI: %v\n\n", err)
		return
	}

	choices, ok := responseMap["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		fmt.Printf("\n\nNo choices found in the response\n\n")
		return
	}
	choice, ok := choices[0].(map[string]interface{})
	if !ok {
		fmt.Printf("\n\nInvalid format in choices\n\n")
		return
	}
	message, ok := choice["message"].(map[string]interface{})
	if !ok {
		fmt.Printf("\n\nInvalid format in message\n\n")
		return
	}
	command, ok := message["content"].(string)
	if !ok {
		fmt.Printf("\n\nInvalid format in content\n\n")
		return
	}

	fmt.Printf("\n\n%s\n\n", command)
}

func getLastTenBashCommand() (string, error) {
	cmd := exec.Command("bash", "-c", "history | tail -n 10")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func sendOpenAIRequest(requestPayload OpenAIRequest, apiKey string) (string, error) {
	payloadBytes, err := json.Marshal(requestPayload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", OpenAIAPIURL, bytes.NewReader(payloadBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
