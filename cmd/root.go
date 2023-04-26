package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "gpt",
	Short: "Simple CLI for Chat GPT",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.simple-chatgpt-cli.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.SetConfigType("yaml")
		viper.AddConfigPath(home)
		viper.SetConfigName(".simple-chatgpt-cli")
	}

	if err := viper.SafeWriteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); !ok {
			log.Fatalf("Error writing config file: %s", err)
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config: ", err)
		os.Exit(1)
	}
}
