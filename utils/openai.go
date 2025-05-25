package utils

import (
	"log"
	"os"
)

func GetBeers(prompt string) string {
	// 1. Get your OpenAI API Key from an environment variable
	//    Set this BEFORE running the script:
	//    Linux/macOS: export OPENAI_API_KEY="sk-YOUR_KEY_HERE"
	//    Windows (cmd): set OPENAI_API_KEY=sk-YOUR_KEY_HERE
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("Error: OPENAI_API_KEY environment variable not set. Please set it before running.")
	}

	return result
}
