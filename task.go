package main

import "time"

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
