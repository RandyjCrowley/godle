package wordle

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type GenerateRequest struct {
	Model   string      `json:"model"`
	Prompt  string      `json:"prompt"`
	Stream  bool        `json:"stream"`
	Options interface{} `json:"options,omitempty"`
}

type GenerateResponse struct {
	Response string `json:"response"`
}

func get_word() string {
	// Set up the request body
	reqBody := GenerateRequest{
		Model:  "llama3",
		Prompt: "You will act as an expert Wordle word generator. I need you to return only a valid 5-letter Wordle word in response to my request. Ensure that your response includes only the word and nothing else—no additional text, explanations, or context. The word should not be from the list of previously used words: USED_ALREADY=[house, horse]. Adhere strictly to providing only a 5-letter word that is not on this list.",
		Stream: false,
	}

	// Marshal the request body to JSON
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalf("Failed to marshal request body: %v", err)
	}

    result := ""
    for {
        // Make the POST request to the Ollama API
        resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(jsonBody))
        if err != nil {
            log.Fatalf("Failed to make request: %v", err)
        }
        defer resp.Body.Close()

        // Read the response body
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatalf("Failed to read response body: %v", err)
        }

        // Unmarshal the response body into a GenerateResponse struct
        var generateResponse GenerateResponse
        err = json.Unmarshal(body, &generateResponse)
        if err != nil {
            log.Fatalf("Failed to unmarshal response body: %v", err)
        }
        if len(generateResponse.Response) == 5 {
            result = generateResponse.Response
            break
        }
    }

	// Return the generated Wordle word as a string
	return result 
}
