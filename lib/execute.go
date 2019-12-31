package lib

import (
	"fmt"
	"os"
	"os/exec"
)

func Execute(command string, args ...string) error {
	fmt.Println("Command:", command, "args:", args)

	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	cmd.Wait()

	return err
}
