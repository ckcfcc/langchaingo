package embeddings_test

import (
	"context"
	"log"

	"github.com/ckcfcc/langchaingo/embeddings"
	"github.com/ckcfcc/langchaingo/llms/openai"
)

func Example() { //nolint:testableexamples
	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Embedder from the given LLM.
	embedder, err := embeddings.NewEmbedder(llm)
	if err != nil {
		log.Fatal(err)
	}

	docs := []string{"doc 1", "another doc"}
	embs, err := embedder.EmbedDocuments(context.Background(), docs)
	if err != nil {
		log.Fatal(err)
	}

	// Consume embs
	_ = embs
}
