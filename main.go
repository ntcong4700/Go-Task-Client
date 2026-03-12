package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	task := Task{
		ID:          1,
		Description: "Learn Go CLI project",
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	// MarshalIndent để in JSON đẹp, dễ nhìn khi test.
	data, err := json.Marshal(task)
	if err != nil {
		fmt.Println("error marshaling task:", err)
		return
	}

	fmt.Println(string(data))
}
