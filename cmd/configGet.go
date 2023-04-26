package cmd

import (
	"github.com/joecottam/simple-chatgpt-cli/config"
	"github.com/spf13/cobra"
)

var configGetCmd = &cobra.Command{
	Use:   "configGet",
	Short: "Get configuration values",
	Run: func(cmd *cobra.Command, args []string) {
		flags := getPresentFlags(cmd)

		if len(flags) == 0 {
			config.PrintConfig(config.GetConfigKeys())
		} else {
			config.PrintConfig(flags)
		}
	},
}

func init() {
	config.SetDefaults()
	for key, desc := range config.GetConfigItems() {
		configGetCmd.Flags().Bool(key, false, desc)
	}
	rootCmd.AddCommand(configGetCmd)
}

func getPresentFlags(cmd *cobra.Command) []string {
	flags := []string{}
	for key := range config.GetConfigItems() {
		value, _ := cmd.Flags().GetBool(key)
		if value {
			flags = append(flags, key)
		}
	}

	return flags
}
