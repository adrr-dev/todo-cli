// Package service contains the service logic
package service

import (
	"fmt"
	"log"
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

	log.Println("new book created", todos[strID])
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
	fmt.Printf("book with id: %s has been deleted\n", id)
	return nil
}

func (s Service) ListTodos() (map[string]*repository.Todo, error) {
	data, err := s.Repo.LoadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s Service) PutTodo(id, content string) (*repository.Todo, error) {
	data, err := s.Repo.LoadAll()
	if err != nil {
		log.Println("something went wrong here")
		return nil, err
	}
	data[id].Content = content
	err = s.Repo.WriteAll(data)
	if err != nil {
		return nil, err
	}
	return data[id], nil
}
