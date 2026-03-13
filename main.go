package main

import (
	"fmt"
	"time"
)

func main() {
	filename := "tasks.json"

	// 1) Load task hiện có từ file.
	tasks, err := LoadTasks(filename)
	if err != nil {
		fmt.Println("error loading tasks:", err)
		return
	}

	fmt.Println("loaded tasks:", len(tasks))

	// 2) Tạo một task mẫu và append vào slice.
	task := Task{
		ID:          1,
		Description: "Learn storage layer",
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, task)

	// 3) Save lại xuống file.
	if err := SaveTasks(filename, tasks); err != nil {
		fmt.Println("error saving tasks:", err)
		return
	}

	fmt.Println("task saved successfully")
}
