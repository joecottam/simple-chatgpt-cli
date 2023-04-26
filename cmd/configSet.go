package cmd

import (
	"fmt"
	"log"

	"github.com/joecottam/simple-chatgpt-cli/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configSetCmd = &cobra.Command{
	Use:   "configSet",
	Short: "Set configuration values",
	Run: func(cmd *cobra.Command, args []string) {
		for key := range config.GetConfigItems() {
			if value, err := cmd.Flags().GetString(key); err == nil && value != "" {
				config.SetConfigValue(key, value)
			}
		}

		if err := viper.WriteConfig(); err != nil {
			log.Fatalf("Error writing config file: %s", err)
		}

		fmt.Println("Configuration has been set.")
	},
}

func init() {
	config.SetDefaults()
	for key, desc := range config.GetConfigItems() {
		configSetCmd.Flags().String(key, "", desc)
	}
	rootCmd.AddCommand(configSetCmd)
}
