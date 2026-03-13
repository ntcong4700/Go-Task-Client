package main

import (
	"encoding/json"
	"errors"
	"os"
)

// function load all the tasks in json file.

func LoadTasks(filename string) ([]Task, error) {
	data, err := os.ReadFile(filename)

	if err != nil {

		if errors.Is(err, os.ErrNotExist) {
			return []Task{}, nil
		}

		return nil, err
	}

	if len(data) == 0 {
		return []Task{}, nil
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func SaveTasks(filename string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
