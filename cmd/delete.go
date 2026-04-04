package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "delete a todo with id",
	Args:  cobra.ExactArgs(1),
	RunE:  deleteTodo,
}

func deleteTodo(cmd *cobra.Command, args []string) error {
	id := args[0]
	fmt.Printf("Removing todo: %s\n", id)
	err := todoService.RemoveTodo(id)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
