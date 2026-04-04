package repository

import (
	"encoding/json"
	"fmt"
	"os"
)

type Repo struct {
	DataFile string
}

func (r Repo) LoadAll() (map[string]*Todo, error) {
	file, err := os.Open(r.DataFile)
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(file)

	defer func() { _ = file.Close() }()

	var data map[string]*Todo
	err = dec.Decode(&data)

	return data, err
}

func (r Repo) WriteAll(data map[string]*Todo) error {
	file, err := os.Create(r.DataFile)
	if err != nil {
		return err
	}

	defer func() { _ = file.Close() }()

	enc := json.NewEncoder(file)
	err = enc.Encode(data)
	return err
}

func (r Repo) GetByID(id string) (*Todo, error) {
	data, err := r.LoadAll()
	if err != nil {
		return nil, err
	}

	todo, ok := data[id]
	if ok {
		return todo, err
	}
	return nil, fmt.Errorf("todo not found or not queried")
}
