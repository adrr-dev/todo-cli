// Package database is where i will manipulate the ORM
package database

import (
	"log"

	"github.com/adrr-dev/todo-cli/internal/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Database struct {
	DataFile string
}

func (d Database) InitializeDB() *gorm.DB {
	DB, err := gorm.Open(sqlite.Open(d.DataFile), &gorm.Config{})
	check(err)

	err = DB.AutoMigrate(&repository.Todo{})
	check(err)
	return DB
}
