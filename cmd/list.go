package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all your todos",
	Args:  cobra.ExactArgs(0),
	RunE:  ShowTodos,
}

func ShowTodos(cmd *cobra.Command, args []string) error {
	data, err := todoService.ListTodos()
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', tabwriter.Debug)

	_, err = fmt.Fprintln(w, "ID\tTodo\tCompleted")
	if err != nil {
		return err
	}
	for _, value := range data {
		_, err = fmt.Fprintf(w, "%d\t%s\t%t\n", value.ID, value.Content, value.Completed)
		if err != nil {
			return err
		}
	}

	err = w.Flush()
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
}
