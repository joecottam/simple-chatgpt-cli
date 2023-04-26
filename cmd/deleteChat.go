package cmd

import (
	"fmt"

	"github.com/joecottam/simple-chatgpt-cli/chat"
	"github.com/joecottam/simple-chatgpt-cli/config"
	"github.com/spf13/cobra"
)

var deleteChatCmd = &cobra.Command{
	Use:   "deleteChat",
	Short: "Delete a saved chat",
	Run: func(cmd *cobra.Command, args []string) {
		historyFileName, err := chat.SelectChat()
		if err != nil {
			fmt.Println("No saved chats found.")
			return
		}
		chat.DeleteChat(historyFileName)
	},
}

func init() {
	config.SetDefaults()
	rootCmd.AddCommand(deleteChatCmd)
}
