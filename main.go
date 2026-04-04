// Package main is the entry point
package main

import (
	"github.com/adrr-dev/todo-cli/cmd"
	"github.com/adrr-dev/todo-cli/internal/repository"
	"github.com/adrr-dev/todo-cli/internal/service"
)

func main() {
	myRepo := &repository.Repo{DataFile: "data.json"}
	myService := &service.Service{Repo: myRepo}
	cmd.SetService(myService)

	cmd.Execute()
}
