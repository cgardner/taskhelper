package cmd

import (
	"fmt"
	"os"

	"github.com/cgardner/taskhelper/types"

	"github.com/spf13/cobra"
)

func NewCommand(template string, props types.Template) *cobra.Command {
	command := &cobra.Command{
		Use:   template,
		Short: fmt.Sprintf("Use the \"%s\" template", template),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				execute("task", props.Report)
				os.Exit(0)
			}
			taskArgs := append([]string{"add"}, props.Add...)
			taskArgs = append(taskArgs, args...)
			execute("task", taskArgs)
		},
	}
	return command
}
