package lib

import (
	"errors"
	"os/exec"

	"github.com/spf13/viper"
)

func GetEditor() (string, error) {
	configLocations := []string{
		"config.editor",
		"editor",
	}

	for _, configName := range configLocations {
		editorName := viper.GetString(configName)
		if editorName == "" {
			continue
		}
		editorCommand, _ := exec.LookPath(editorName)
		if editorCommand != "" {
			return editorCommand, nil
		}
	}

	return "", errors.New("Failed to find an Editor")

}
