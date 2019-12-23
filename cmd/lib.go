package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execute(command string, parameters []string) error {
	fmt.Println("Running `" + command + "` with \"" + strings.Join(parameters, " ") + "\"")
	cmd := getCmd(command, parameters)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	cmd.Wait()
	return err
}

func getCmd(command string, parameters []string) *exec.Cmd {
	if len(parameters) == 1 {
		var parameter = parameters[0]
		return exec.Command(command, parameter)
	}
	return exec.Command(command, parameters...)
}
