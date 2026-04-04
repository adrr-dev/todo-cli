package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [id] [content]",
	Short: "edit an existing todo",
	Args:  cobra.ExactArgs(0),
	RunE:  editTodo,
}

func editTodo(cmd *cobra.Command, args []string) error {
	id, err := cmd.Flags().GetInt("id")
	if err != nil {
		return err
	}
	content, err := cmd.Flags().GetString("content")
	if err != nil {
		return err
	}

	err = todoService.EditTodo(id, content)
	if err != nil {
		return err
	}
	fmt.Printf("new todo created\nid:'%d'\ncontent:%s\n", id, content)
	return nil
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().IntP("id", "i", 0, "id to edit")
	editCmd.Flags().StringP("content", "c", "empty", "content of todo")

	if err := editCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
	}
	if err := editCmd.MarkFlagRequired("content"); err != nil {
		fmt.Println(err)
	}
}
