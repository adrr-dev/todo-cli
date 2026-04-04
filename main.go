// Package main is the entry point
package main

import (
	"github.com/adrr-dev/todo-cli/cmd"
	"github.com/adrr-dev/todo-cli/internal/database"
	"github.com/adrr-dev/todo-cli/internal/service"
)

func main() {
	dataFile := "data.db"

	myDB := database.Database{DataFile: dataFile}
	db := myDB.InitializeDB()

	myService := &service.Service{DB: db}
	cmd.SetService(myService)

	cmd.Execute()
}
