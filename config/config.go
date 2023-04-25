package config

import (
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

var configItems = map[string]string{
	"openAiKey":    "Open AI API key",
	"chatUserName": "The name of the user that will appear in the chat",
	"chatsDir":     "The directory to save chat history files to",
	"defaultModel": "The default model to use",
}

func SetDefaults() {
	viper.SetDefault("openAiKey", "")
	viper.SetDefault("chatUserName", "User")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = ""
	}
	viper.SetDefault("chatsDir", fmt.Sprintf("%v/simple-chatgpt-cli-chats", homeDir))
	viper.SetDefault("defaultModel", openai.GPT3Dot5Turbo)
}

func GetConfigItems() map[string]string {
	return configItems
}

func GetSetConfigItems() map[string]string {
	items := map[string]string{}

	for key, value := range configItems {
		if viper.IsSet(key) {
			items[key] = value
		}
	}
	return items
}

func GetConfigValue(key string) string {
	return viper.GetString(key)
}

func GetConfigKeys() []string {
	keys := []string{}
	for key := range configItems {
		keys = append(keys, key)
	}
	return keys
}

func SetConfigValue(key string, value string) {
	viper.Set(key, value)
}

func PrintConfig(keys []string) {
	for _, key := range keys {
		str := key + ": "
		if viper.InConfig(key) {
			str = str + GetConfigValue(key)
		} else {
			str = str + "unset (default: " + GetConfigValue(key) + ")"
		}
		fmt.Println(str)
	}
}
