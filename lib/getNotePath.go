package lib

import (
	"os"

	"github.com/spf13/viper"
)

func GetNotePath() (string, error) {
	// Get Configured path for task notes.  Default to ~/.task
	noteRoot := viper.GetString("config.notes.path")
	notePath := os.ExpandEnv(noteRoot)

	// pathInfo, err := os.Stat(notePath)
	_, err := os.Stat(notePath)
	if err != nil {
		os.MkdirAll(notePath, os.ModePerm)
	}

	return notePath, nil
}
