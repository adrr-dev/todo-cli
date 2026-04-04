package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "toggle completion of a todo with it's id",
	Args:  cobra.ExactArgs(1),
	RunE:  completeTodo,
}

func completeTodo(cmd *cobra.Command, args []string) error {
	id := args[0]
	err := todoService.CompleteTodo(id)
	if err != nil {
		return err
	}
	fmt.Printf("todo with id %s complete status changed\n", id)
	return nil
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
