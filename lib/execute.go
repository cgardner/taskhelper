package lib

import (
	"os"
	"os/exec"
)

func Execute(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	cmd.Wait()

	return err
}
