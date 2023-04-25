package cmd

import (
	"github.com/joecottam/simple-chatgpt-cli/config"
	"github.com/spf13/cobra"
)

// configGetCmd represents the configGet command
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
	for key, desc := range config.GetConfigItems() {
		configGetCmd.Flags().String(key, "", desc)
		configGetCmd.Flags().Lookup(key).NoOptDefVal = "true"
	}
	rootCmd.AddCommand(configGetCmd)
}

func getPresentFlags(cmd *cobra.Command) []string {
	flags := []string{}
	for key := range config.GetConfigItems() {
		value, _ := cmd.Flags().GetString(key)
		if value != "" {
			flags = append(flags, key)
		}
	}

	return flags
}
