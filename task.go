package main

import (
	"errors"
	"strings"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

const (
	StatusTodo      = "todo"
	StatusInProcess = "in-progress"
	StatusDone      = "done"
)

func IsValidStatus(status string) bool {
	return status == StatusTodo || status == StatusInProcess || status == StatusDone
}

func NewTask(id int, description string, status string) (Task, error) {
	description = strings.TrimSpace(description)
	if description == "" {
		return Task{}, errors.New("description can be not empty")
	}

	now := time.Now()
	task := Task{
		ID:          id,
		Description: description,
		Status:      status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return task, nil
}

func NextTaskID(tasks []Task) int {
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}

func AddTask(tasks []Task, description string) ([]Task, Task, error) {
	id := NextTaskID(tasks)
	task, err := NewTask(id, description, StatusTodo)
	if err != nil {
		return nil, Task{}, err
	}

	tasks = append(tasks, task)

	return tasks, task, nil
}

func UpdateTask(tasks []Task, id int, description string) ([]Task, error) {
	description = strings.TrimSpace(description)

	if description == "" {
		return nil, errors.New("description can be not empty")
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			return tasks, nil
		}
	}

	return nil, errors.New("task not found")
}

func DeleteTask(tasks []Task, id int) ([]Task, error) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return tasks, nil
		}
	}
	return nil, errors.New("task not found")
}
