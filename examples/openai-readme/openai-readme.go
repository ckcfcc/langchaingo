package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ckcfcc/langchaingo/llms"
	"github.com/ckcfcc/langchaingo/llms/openai"
)

func main() {
	ctx := context.Background()
	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}
	prompt := "What would be a good company name for a company that makes colorful socks?"
	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(completion)
}
