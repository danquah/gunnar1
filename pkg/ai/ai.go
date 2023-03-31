package ai

import (
	"context"
	"errors"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type OpenAIClient struct {
	client   *openai.Client
	language string
}

const defaultPrompt = "The following is a conversation with an AI assistant. The assistant is helpful, creative, clever, and very friendly. Reply to in %s, here is the question: %s"

func (c *OpenAIClient) Configure(apiToken string, language string) error {
	client := openai.NewClient(apiToken)
	if client == nil {
		return errors.New("failed to create OpenAI client")
	}
	c.language = language
	c.client = client

	return nil
}

func (c *OpenAIClient) Complete(ctx context.Context, prompt string) (string, error) {
	if c.client == nil {
		return "", errors.New("OpenAI client is not configured")
	}

	resp, err := c.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: fmt.Sprintf(defaultPrompt, c.language, prompt),
			},
		},
	})
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
