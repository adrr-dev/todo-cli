// Package cmd contains the cmds
package cmd

import (
	"os"

	"github.com/adrr-dev/todo-cli/internal/repository"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "a simple todo cli",
	Long:  `a todo cli that can add, delete, list, update, and complete your todo`,
}

type TodoService interface {
	CreateTodo(content string) error
	FetchTodo(id string) (*repository.Todo, error)
	RemoveTodo(id string) error
	ListTodos() (map[string]*repository.Todo, error)
	EditTodo(id, content string) error
	CompleteTodo(id string) error
}

var todoService TodoService

// SetService allows your main.go to "plug in" the implementation
func SetService(s TodoService) {
	todoService = s
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todo-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
