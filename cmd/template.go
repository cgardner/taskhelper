package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/cgardner/taskhelper/types"

	"github.com/spf13/cobra"
)

func NewCommand(template string, props types.Template) *cobra.Command {
	command := &cobra.Command{
		Use:   fmt.Sprintf("%s [description] [...args]", template),
		Short: fmt.Sprintf("Use the \"%s\" template", template),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				displayReport(props.Report)
				return
			}

			task := props.Add
			task.Description = args[0]

			err := saveTask(task)
			if err != nil {
				fmt.Println(err.Error())
			}

		},
	}
	return command
}

func saveTask(newTask types.TaskTemplate) error {
	tasks, err := json.Marshal(newTask)
	if err != nil {
		return err
	}

	cmd := exec.Command("task", "import", "-")
	cmd.Stdin = bytes.NewBuffer(tasks)
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func displayReport(report types.Report) {
	params := []string{}

	if report.Filters.Project != "" {
		params = append(params, fmt.Sprintf("project:\"%s\"", report.Filters.Project))
	}

	if len(report.Filters.Tags) != 0 {
		for _, tag := range report.Filters.Tags {
			params = append(params, fmt.Sprintf("+%s", tag))
		}
	}

	if report.Report != "" {
		params = append(params, report.Report)
	}

	cmd := exec.Command("task", params...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	cmd.Wait()
	if err != nil {
		fmt.Println(err.Error())
	}
}
