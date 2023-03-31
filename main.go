package main

import (
	"context"
	"fmt"
	"os"

	"github.com/danquah/gunnar1/pkg/ai"
)

func main() {
	// Get OPEN_AI_TOKEN from environment
	apiToken := os.Getenv("OPEN_AI_TOKEN")
	if apiToken == "" {
		panic("OPEN_AI_TOKEN is not set")
	}

	if len(os.Args) == 1 {
		panic("no prompt provided")
	}
	ctx := context.Background()

	prompt := os.Args[1]
	client := &ai.OpenAIClient{}
	err := client.Configure(apiToken, "english")
	if err != nil {
		panic(err)
	}
	output, err := client.Complete(ctx, prompt)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
