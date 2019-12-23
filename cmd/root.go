package cmd

import (
	"fmt"
	"os"

	"github.com/cgardner/taskhelper/types"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "taskhelper",
	Short: "taskwarrior helper",
	Long:  "This is a helper function for taskwarrior that allows you to set up task templates and generate reports",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	viper.SetConfigName("taskhelper")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/taskhelper")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	var settings map[string]types.Template
	viper.Unmarshal(&settings)

	for name, props := range settings {
		command := NewCommand(name, props)
		rootCmd.AddCommand(command)
	}
}

// Configuration Initialization
func initConfig() {
	viper.SetConfigName(".taskhelper")
	viper.AddConfigPath("$HOME/.config/.taskhelper")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// Execute the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
