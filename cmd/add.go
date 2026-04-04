package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [todo]",
	Short: "Add a new todo to your list",
	Args:  cobra.ExactArgs(1),
	RunE:  addTodo,
}

func addTodo(cmd *cobra.Command, args []string) error {
	todo := args[0]
	err := todoService.CreateTodo(todo)
	if err != nil {
		return err
	}
	fmt.Printf("new todo created '%s'\n", todo)
	return nil
}

func init() {
	rootCmd.AddCommand(addCmd)
}
