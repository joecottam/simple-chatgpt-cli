package chat

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joecottam/simple-chatgpt-cli/config"
	"github.com/manifoldco/promptui"
	"github.com/sashabaranov/go-openai"
)

type History struct {
	Messages []openai.ChatCompletionMessage `json:"messages"`
}

func (h *History) save(filepath string) {
	content, err := json.Marshal(h)
	if err != nil {
		fmt.Println("Error marshalling json: ", err)
	}

	err = os.WriteFile(filepath, content, 0644)

	if err != nil {
		fmt.Println("Error writing to file: ", err)
	}
}

func (h *History) load(filepath string) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	err = json.Unmarshal(content, &h)
	if err != nil {
		log.Fatal("Error unmarshalling json: ", err)
	}
}

func SelectChat() (string, error) {
	chatFileNames := []string{}
	files, _ := os.ReadDir(config.GetConfigValue("chatsDir"))

	if len(files) == 0 {
		return "", fmt.Errorf("no saved chats found")
	}

	for _, file := range files {
		chatFileNames = append(chatFileNames, file.Name())
	}

	prompt := promptui.Select{
		Label: "Select chat",
		Items: chatFileNames,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

func DeleteChat(historyFilename string) {
	historyFilePath := fmt.Sprintf("%v/%v", config.GetConfigValue("chatsDir"), historyFilename)

	err := os.Remove(historyFilePath)
	if err != nil {
		log.Fatal("Error deleting file: ", err)
	}
	fmt.Printf("Chat successfully deleted: %v\n", historyFilePath)
}
