package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Gemini API endpoint for generateContent
const geminiAPIURL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"

// Request payload structure for Gemini API
type GeminiRequest struct {
	Contents []Content `json:"contents"`
}

type Content struct {
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

// Response payload structure from Gemini API (simplified for this example)
type GeminiResponse struct {
	Candidates []Candidate `json:"candidates"`
	// Add other fields like promptFeedback, usageMetadata if needed
}

type Candidate struct {
	Content Content `json:"content"`
	// Add other fields like safetyRatings, finishReason if needed
}

func GetBeers(prompt string) string {
	// Get Gemini API Key from Environment Variable
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("Error: OPENAI_API_KEY environment variable not set. Please set it before running.")
	}

	// Construct the request payload
	requestPayload := GeminiRequest{
		Contents: []Content{
			{
				Parts: []Part{
					{
						Text: prompt,
					},
				},
			},
		},
	}

	// Marshal the request payload to JSON
	jsonPayload, err := json.Marshal(requestPayload)
	if err != nil {
		log.Fatalf("Error marshaling request payload: %v", err)
	}

	// Construct the full URL with the API key
	fullURL := fmt.Sprintf("%s?key=%s", geminiAPIURL, apiKey)

	// Create an HTTP client
	client := &http.Client{}

	// Create the HTTP request
	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request
	fmt.Println("Sending request to Gemini API...")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close() // Ensure the response body is closed

	// Read the response body

	responseBody := &GeminiResponse{}
	derr := json.NewDecoder(resp.Body).Decode(responseBody)
	if derr != nil {
		panic(derr)
	}

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Gemini API returned an error: Status %d, Response: %s", resp.StatusCode, responseBody)
	}

	beers := responseBody.Candidates[0].Content.Parts[0].Text

	return beers
}
