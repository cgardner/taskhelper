package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"os/exec"
)

// AddAction Flag to determine whether the user is adding a task with this template
var AddAction bool

var rootCmd = &cobra.Command{
	Use:   "taskhelper",
	Short: "taskwarrior helper",
	Long:  "This is a helper function for taskwarrior that allows you to set up task templates and generate reports",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var template string = args[0]

		if AddAction {
			add(template, shift(args))
			return
		}
		report(template)
	},
}

func report(template string) {
	var parameters = []string{"+" + template, ""}
	err := execute("task", parameters)
	fmt.Printf("Report task finished with :%v", err)
	fmt.Println("Displaying " + template + " report")
}

func add(template string, task []string) {
	parameters := append([]string{"add", "+" + template}, task...)

	fmt.Println("Adding a task using \"" + template + "\"")
	err := execute("task", parameters)
	fmt.Printf("Command finished with error: %v", err)
	return
}

func execute(command string, parameters []string) error {
	cmd := getCmd(command, parameters)
	out, err := cmd.Output()
	fmt.Println(string(out))
	return err
}

func print(readCloser io.ReadCloser) {
	r := bufio.NewReader(readCloser)
	line, _, err := r.ReadLine()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(string(line))
}

func getCmd(command string, parameters []string) *exec.Cmd {
	if len(parameters) == 1 {
		var parameter = parameters[0]
		return exec.Command(command, parameter)
	}
	return exec.Command(command, parameters...)
}

// Shift an element
func shift(s []string) []string {
	var shifted = s[1:]
	return shifted
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&AddAction, "add", "a", false, "Add a new task with the template")
}

// Execute the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
