// Package cmd contains the cmds
package cmd

import (
	"os"

	"github.com/adrr-dev/todo-cli/internal/repository"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "a simple todo cli",
	Long:  `a todo cli that can add, delete, list, update, and complete your todo`,
}

type TodoService interface {
	CreateTodo(content string) error
	RemoveTodo(id int) error
	ListTodos() ([]*repository.Todo, error)
	EditTodo(id int, content string) error
	CompleteTodo(id int) error
}

var todoService TodoService

func SetService(s TodoService) {
	todoService = s
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
