package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	chatcount := 10

	for chatcount <= 10 {
		//load .env file
		godotenv.Load()

		apikey := os.Getenv("api_key")
		client := openai.NewClient(apikey)

		//get the user input
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println()
		fmt.Print("User: ")
		scanner.Scan()
		input := scanner.Text()
		fmt.Println()

		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: input,
					},
				},
			},
		)

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			return
		}

		fmt.Println(resp.Choices[0].Message.Content)
		chatcount--

	}
}
