// Package main is the entry point
package main

import (
	"github.com/adrr-dev/todo-cli/cmd"
	"github.com/adrr-dev/todo-cli/internal/repository"
	"github.com/adrr-dev/todo-cli/internal/service"
)

func main() {
	dataFile := "data_test.json"
	myRepo := &repository.Repo{DataFile: dataFile}
	myService := &service.Service{Repo: myRepo}
	cmd.SetService(myService)

	cmd.Execute()
}
