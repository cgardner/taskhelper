package cmd

import (
	"fmt"
	"os"

	"github.com/cgardner/taskhelper/types"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string

	rootCmd = &cobra.Command{
		Version: "0.1.0",
		Use:     "taskhelper",
		Short:   "taskwarrior helper",
		Long:    "This is a helper function for taskwarrior that allows you to set up task templates and generate reports",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}

			subCmd, _, _ := cmd.Traverse(args)
			subCmd.Execute()
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			var settings map[string]types.Template
			viper.Unmarshal(&settings)

			for name, props := range settings {
				command := NewCommand(name, props)
				cmd.AddCommand(command)
			}
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file (default is $HOME/.config/taskhelper.yaml)")
}

// Configuration Initialization
func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("taskhelper")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("$HOME/.config/taskhelper")
	}
	viper.AutomaticEnv()

	viper.ReadInConfig()
}

// Execute the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
