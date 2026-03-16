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
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

func IsValidStatus(status string) bool {
	return status == StatusTodo || status == StatusInProgress || status == StatusDone
}

func NewTask(id int, description string) (Task, error) {
	description = strings.TrimSpace(description)
	if description == "" {
		return Task{}, errors.New("description cannot be empty")
	}

	now := time.Now()

	task := Task{
		ID:          id,
		Description: description,
		Status:      StatusTodo,
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

	task, err := NewTask(id, description)
	if err != nil {
		return nil, Task{}, err
	}

	tasks = append(tasks, task)
	return tasks, task, nil
}

func UpdateTask(tasks []Task, id int, newDescription string) ([]Task, error) {
	newDescription = strings.TrimSpace(newDescription)
	if newDescription == "" {
		return nil, errors.New("new description cannot be empty")
	}

	for i := range tasks {
		if tasks[i].ID == id {
			if tasks[i].Description == newDescription {
				return nil, errors.New("new description is the same as the current one")
			}

			tasks[i].Description = newDescription
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

func MarkTaskStatus(tasks []Task, id int, status string) ([]Task, error) {
	if !IsValidStatus(status) {
		return nil, errors.New("invalid status")
	}

	for i := range tasks {
		if tasks[i].ID == id {
			if tasks[i].Status == status {
				return nil, errors.New("task already has this status")
			}

			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			return tasks, nil
		}
	}

	return nil, errors.New("task not found")
}

func ListTasks(tasks []Task, status string) ([]Task, error) {
	if status == "" {
		return tasks, nil
	}

	if !IsValidStatus(status) {
		return nil, errors.New("invalid status filter")
	}

	var filteredTasks []Task
	for _, task := range tasks {
		if task.Status == status {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks, nil
}
