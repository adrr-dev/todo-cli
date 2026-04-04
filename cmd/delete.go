package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "delete a todo with id",
	Args:  cobra.ExactArgs(0),
	RunE:  deleteTodo,
}

func deleteTodo(cmd *cobra.Command, args []string) error {
	id, err := cmd.Flags().GetInt("id")
	if err != nil {
		return err
	}
	fmt.Printf("Removing todo: %d\n", id)
	err = todoService.RemoveTodo(id)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntP("id", "i", -1, "id to edit")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
	}
}
