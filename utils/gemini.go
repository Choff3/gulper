package utils

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"google.golang.org/genai"
)

const basePrompt = "Get a list of all the beers on this menu and return each of them in an array of JSONs with each attribute as a key. " +
	"The keys are: name, brewery, style, abv, and price. " +
	"Store any other information in a separate key based on the label in the menu or description if no label is provided."

func GetBeersHTML(url string, addPrompt string) string {
	// Get Gemini API Key from Environment Variable
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("Error: GEMINI_API_KEY environment variable not set. Please set it before running.")
	}

	prompt := fmt.Sprintf("%s %s %s", basePrompt, addPrompt, url)

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	return result.Text()
}

func GetBeersPDF(url string, addPrompt string) string {

	ctx := context.Background()
	client, _ := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})

	prompt := fmt.Sprintf("%s %s", basePrompt, addPrompt)

	pdfBytes := getPDF(url)

	parts := []*genai.Part{
		&genai.Part{
			InlineData: &genai.Blob{
				MIMEType: "application/pdf",
				Data:     pdfBytes,
			},
		},
		genai.NewPartFromText(prompt),
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
	}

	result, _ := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		contents,
		nil,
	)

	return result.Text()
}

func getPDF(url string) []byte {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", GetUserAgent())

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	pdfBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return pdfBytes
}

func GetUserAgent() string {
	return "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:138.0) Gecko/20100101 Firefox/138.0"
}
