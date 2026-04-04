// Package service contains the service logic
package service

import (
	"fmt"
	"strconv"

	"github.com/adrr-dev/todo-cli/internal/repository"
)

type TodoRepo interface {
	LoadAll() (map[string]*repository.Todo, error)
	WriteAll(map[string]*repository.Todo) error
	GetByID(string) (*repository.Todo, error)
}

type Service struct {
	Repo TodoRepo
}

func (s Service) CreateTodo(content string) error {
	todos, err := s.Repo.LoadAll()
	if err != nil {
		return err
	}
	newID := 1
	var strID string
	for {
		strID = strconv.Itoa(newID)
		_, ok := todos[strID]
		if !ok {
			todos[strID] = &repository.Todo{ID: strID, Content: content, Completed: false}
			break
		}
		newID++
	}

	err = s.Repo.WriteAll(todos)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) FetchTodo(id string) (*repository.Todo, error) {
	todo, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (s Service) RemoveTodo(id string) error {
	data, err := s.Repo.LoadAll()
	if err != nil {
		return err
	}

	delete(data, id)

	err = s.Repo.WriteAll(data)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) ListTodos() (map[string]*repository.Todo, error) {
	data, err := s.Repo.LoadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s Service) EditTodo(id, content string) error {
	data, err := s.Repo.LoadAll()
	if err != nil {
		return err
	}
	data[id].Content = content
	err = s.Repo.WriteAll(data)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) CompleteTodo(id string) error {
	data, err := s.Repo.LoadAll()
	if err != nil {
		return err
	}

	value, ok := data[id]
	if ok {
		value.Completed = !value.Completed
	} else {
		return fmt.Errorf("todo not found: %e", err)
	}

	err = s.Repo.WriteAll(data)
	if err != nil {
		return err
	}

	return nil
}
