package cmd

import (
	"fmt"

	"github.com/joecottam/simple-chatgpt-cli/chat"
	"github.com/joecottam/simple-chatgpt-cli/config"
	"github.com/spf13/cobra"
)

var loadChatCmd = &cobra.Command{
	Use:   "loadChat",
	Short: "Load a saved chat",
	Run: func(cmd *cobra.Command, args []string) {
		c := chat.Chat{History: chat.History{}}
		historyFileName, err := chat.SelectChat()
		if err != nil {
			fmt.Println("No saved chats found.")
			fmt.Println("Starting new chat")
		} else {
			c.LoadHistory(historyFileName)
		}

		c.Start()
	},
}

func init() {
	config.SetDefaults()
	rootCmd.AddCommand(loadChatCmd)
}
