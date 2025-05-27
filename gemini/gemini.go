package gemini

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Choff3/gulper/utils"
	"google.golang.org/genai"
)

const basePrompt = "Get a list of all the beers on this menu and return each of them in an array of JSONs with each attribute as a key. " +
	"The keys are: name, brewery, style, abv, and price."

func GetMenuHTML(url string, addPrompt string) string {
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
		getConfig(),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.Text())

	return result.Text()
}

func GetMenuPDF(url string, addPrompt string) string {

	ctx := context.Background()
	client, _ := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})

	prompt := fmt.Sprintf("%s %s", basePrompt, addPrompt)

	fmt.Println("Parsing menu from", url)
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

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		contents,
		getConfig(),
	)
	if err != nil {
		log.Fatal(err)
	}

	return result.Text()
}

func getPDF(url string) []byte {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", utils.GetUserAgent())

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

func getConfig() *genai.GenerateContentConfig {
	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema: &genai.Schema{
			Type: genai.TypeArray,
			Items: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"name":    {Type: genai.TypeString},
					"brewery": {Type: genai.TypeString},
					"style":   {Type: genai.TypeString},
					"abv":     {Type: genai.TypeNumber},
					"price":   {Type: genai.TypeNumber},
				},
				PropertyOrdering: []string{"name", "brewery", "style", "abv", "price"},
			},
		},
	}
	return config
}
