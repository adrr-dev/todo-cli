// Package service contains the service logic
package service

import (
	"github.com/adrr-dev/todo-cli/internal/repository"

	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func (s Service) FetchTodo(id int) (*repository.Todo, error) {
	var todo repository.Todo

	result := s.DB.First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}

func (s Service) CreateTodo(content string) error {
	todo := &repository.Todo{Content: content, Completed: false}
	result := s.DB.Create(todo)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s Service) RemoveTodo(id int) error {
	result := s.DB.Delete(&repository.Todo{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s Service) ListTodos() ([]*repository.Todo, error) {
	var todos []*repository.Todo
	result := s.DB.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return todos, result.Error
}

func (s Service) EditTodo(id int, content string) error {
	result := s.DB.Model(&repository.Todo{}).Where("id = ?", id).Update("content", content)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s Service) CompleteTodo(id int) error {
	todo, err := s.FetchTodo(id)
	if err != nil {
		return err
	}

	completed := todo.Completed
	result := s.DB.Model(&repository.Todo{}).Where("id = ?", id).Update("completed", !completed)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
