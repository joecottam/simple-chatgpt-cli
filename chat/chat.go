package chat

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/joecottam/simple-chatgpt-cli/config"
	"github.com/sashabaranov/go-openai"
)

type Chat struct {
	History         History
	SystemMessage   string
	Chatting        bool
	HistoryFileName string
}

func (c *Chat) Start() {
	c.Chatting = true

	c.getSystemMessage()
	c.getUserMessage()

	for c.Chatting {
		c.getAssistantMessage()
		c.getUserMessage()
	}

	fmt.Printf("Chat history saved to %v\n", c.historyFilePath())
}

func (c *Chat) getSystemMessage() {
	if c.SystemMessage != "" {
		fmt.Print(color.RedString("\nSystem: "))
		fmt.Print(c.SystemMessage)
		fmt.Print("\n\n")
		systemChatMessage := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: c.SystemMessage,
		}
		c.History.Messages = append(c.History.Messages, systemChatMessage)
	}
}

func (c *Chat) getAssistantMessage() {
	client := openai.NewClient(config.GetConfigValue("openAiKey"))
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model:    c.History.Model,
		Messages: c.History.Messages,
		Stream:   true,
	}
	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		log.Fatalf("ChatCompletionStream error: %v", err)
	}
	defer stream.Close()

	var messageBody string

	fmt.Print("\n")
	fmt.Print(color.YellowString("Assistant: "))

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Print("\n\n")
			message := openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleAssistant,
				Content: messageBody,
			}
			c.History.Messages = append(c.History.Messages, message)
			c.History.save(c.historyFilePath())
			return
		}

		if err != nil {
			fmt.Printf("stream error: %v", err)
		}

		messageBody = messageBody + response.Choices[0].Delta.Content
		fmt.Printf(response.Choices[0].Delta.Content)
	}
}

func (c *Chat) getUserMessage() {
	fmt.Print(color.GreenString(config.GetConfigValue("chatUserName") + ": "))
	reader := bufio.NewReader(os.Stdin)

	content, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Reader error: ", err)
	}

	if content == "exit\n" {
		c.Chatting = false
	}

	c.History.Messages = append(c.History.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: content,
	})
}

func (c *Chat) historyFilePath() string {
	return fmt.Sprintf("%v/%v", config.GetConfigValue("chatsDir"), c.HistoryFileName)
}

func (c *Chat) LoadHistory(historyFileName string) {
	c.HistoryFileName = historyFileName
	historyFilePath := c.historyFilePath()
	fmt.Println("Loading chat history from: ", historyFilePath)
	c.History.load(historyFilePath)

	for _, message := range c.History.Messages {
		if message.Role == openai.ChatMessageRoleSystem {
			fmt.Print(color.RedString("\nSystem: "))
			fmt.Print(message.Content)
			fmt.Print("\n")
		}
		if message.Role == openai.ChatMessageRoleAssistant {
			fmt.Print(color.YellowString("\nAssistant: "))
			fmt.Print(message.Content)
			fmt.Print("\n\n")
		}
		if message.Role == openai.ChatMessageRoleUser {
			fmt.Print(color.GreenString("\n" + config.GetConfigValue("chatUserName") + ": "))
			fmt.Print(message.Content)
		}
	}
}
