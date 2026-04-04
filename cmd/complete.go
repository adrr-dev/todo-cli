package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "toggle completion of a todo with it's id",
	Args:  cobra.ExactArgs(0),
	RunE:  completeTodo,
}

func completeTodo(cmd *cobra.Command, args []string) error {
	id, err := cmd.Flags().GetInt("id")
	if err != nil {
		return err
	}
	err = todoService.CompleteTodo(id)
	if err != nil {
		return err
	}
	fmt.Printf("todo with id %d complete status changed\n", id)
	return nil
}

func init() {
	rootCmd.AddCommand(completeCmd)

	completeCmd.Flags().IntP("id", "i", -1, "id to edit")

	if err := completeCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
	}
}
