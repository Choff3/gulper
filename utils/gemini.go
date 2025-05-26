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

const basePrompt = "Get a list of all the beers on this menu and return each of them in an array of JSONs with each column as a key. " +
	"The keys are: name, brewery, style, abv, description, and price."

func GetBeersHTML(url string) string {
	// Get Gemini API Key from Environment Variable
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("Error: GEMINI_API_KEY environment variable not set. Please set it before running.")
	}

	prompt := fmt.Sprintf("%s %s", basePrompt, url)

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

func GetBeersPDF(url string) string {

	ctx := context.Background()
	client, _ := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})

	pdfBytes := getPDF(url)

	parts := []*genai.Part{
		&genai.Part{
			InlineData: &genai.Blob{
				MIMEType: "application/pdf",
				Data:     pdfBytes,
			},
		},
		genai.NewPartFromText("Summarize this document."),
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
