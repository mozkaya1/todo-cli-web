package storage

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	Filename string
}

func (s *Storage[T]) SaveFile(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.Filename, fileData, 0644)
}

func (s *Storage[T]) LoadFile(data *T) error {
	filedata, err := os.ReadFile("todo.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(filedata, data)
}
