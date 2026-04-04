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
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		fmt.Printf("Adding todo: %s\n", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
