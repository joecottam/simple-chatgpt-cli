package cmd

import (
	"fmt"
	"time"

	"github.com/joecottam/simple-chatgpt-cli/chat"
	"github.com/joecottam/simple-chatgpt-cli/config"
	"github.com/spf13/cobra"
)

var historyFileName = fmt.Sprintf("chat_%v.json", time.Now().Unix())
var systemMessage string
var loadHistory string

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat!",
	Run: func(cmd *cobra.Command, args []string) {
		chat := chat.Chat{
			History:         chat.History{},
			SystemMessage:   systemMessage,
			HistoryFileName: historyFileName,
		}

		chat.Start()
	},
}

func init() {
	config.SetDefaults()
	chatCmd.Flags().StringVarP(&systemMessage, "systemMessage", "", "", "The system message i.e. the prompt given to the AI")
	chatCmd.Flags().StringVarP(&historyFileName, "historyFileName", "", historyFileName, "The file name to save the chat history to")
	chatCmd.Flags().StringVarP(&loadHistory, "loadHistory", "", "", "The file name to load the chat history from")
	rootCmd.AddCommand(chatCmd)
}
