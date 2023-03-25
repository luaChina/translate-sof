package service

import (
	"context"
	"github.com/luaChina/translate-sof/config"
	openai "github.com/sashabaranov/go-openai"
	"net/http"
)

// SendChatMessage sends a message to the OpenAI API and returns the response
func SendChatMessage(ctx context.Context, message string) (string, error) {
	gptConfig := openai.DefaultConfig(config.SecretConfig.ApiKey)
	//proxyUrl, err := url.Parse(config.SecretConfig.ApiProxy)
	//if err != nil {
	//	panic(err)
	//}
	transport := &http.Transport{
		//Proxy: http.ProxyURL(proxyUrl),
	}
	gptConfig.HTTPClient = &http.Client{
		Transport: transport,
	}
	client := openai.NewClientWithConfig(gptConfig)
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
