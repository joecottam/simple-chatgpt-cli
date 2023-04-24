package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var configItems = map[string]string{
	"openAiKey":    "Open AI API key",
	"chatUserName": "The name of the user that will appear in the chat",
	"chatsDir":     "The directory to save chat history files to",
}

func SetDefaults() {
	viper.SetDefault("openAiKey", "")
	viper.SetDefault("chatUserName", "User")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = ""
	}
	viper.SetDefault("chatsDir", fmt.Sprintf("%v/simple-chatgpt-cli-chats", homeDir))
}

func GetConfigItems() map[string]string {
	return configItems
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
