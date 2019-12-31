package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/cgardner/taskhelper/lib"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit notes for a task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task, err := lib.GetNote(args[0])
		if err != nil {
			log.Fatal(err)
		}

		noteRoot, err := lib.GetNotePath()
		if err != nil {
			log.Fatal(err)
		}

		fileName := fmt.Sprintf("%s/%s.md", noteRoot, task.Uuid())

		f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}

		editor, err := lib.GetEditor()
		if err != nil {
			log.Fatal(err)
		}

		err = lib.Execute(editor, fileName)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
