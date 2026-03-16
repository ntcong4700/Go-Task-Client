package main

import (
	"errors"
	"fmt"
	"strconv"
)

func PrintUsage() {
	fmt.Println("Task Tracker CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println(`  go run . add "task description"`)
	fmt.Println(`  go run . update <id> "new description"`)
	fmt.Println(`  go run . delete <id>`)
	fmt.Println(`  go run . mark-in-progress <id>`)
	fmt.Println(`  go run . mark-done <id>`)
	fmt.Println(`  go run . list`)
	fmt.Println(`  go run . list done`)
	fmt.Println(`  go run . list todo`)
	fmt.Println(`  go run . list in-progress`)
}

func ParseID(value string) (int, error) {
	id, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Invalid ID: %s", value))

	}
	if id <= 0 {
		return 0, errors.New(fmt.Sprintf("Invalid ID: %d", id))
	}
	return id, nil
}

func PrintTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for _, task := range tasks {
		fmt.Println("--------------------------------------------------")
		fmt.Printf("ID          : %d\n", task.ID)
		fmt.Printf("Description : %s\n", task.Description)
		fmt.Printf("Status      : %s\n", task.Status)
		fmt.Printf("Created At  : %s\n", task.CreatedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("Updated At  : %s\n", task.UpdatedAt.Format("2006-01-02 15:04:05"))
	}
	fmt.Println("--------------------------------------------------")
}

// -----------------------handle-----------------------------
// handleAdd xử lý command: add "description"
func handleAdd(tasks []Task, args []string) {
	if len(args) != 3 {
		fmt.Println(`usage: go run . add "task description"`)
		osExit1()
	}

	updatedTasks, task, err := AddTask(tasks, args[2])
	if err != nil {
		fmt.Println("error adding task:", err)
		osExit1()
	}

	if err := SaveTasks(taskFile, updatedTasks); err != nil {
		fmt.Println("error saving tasks:", err)
		osExit1()
	}

	fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
}

// handleUpdate xử lý command: update <id> "new description"
func handleUpdate(tasks []Task, args []string) {
	if len(args) != 4 {
		fmt.Println(`usage: go run . update <id> "new description"`)
		osExit1()
	}

	id, err := ParseID(args[2])
	if err != nil {
		fmt.Println("error:", err)
		osExit1()
	}

	updatedTasks, err := UpdateTask(tasks, id, args[3])
	if err != nil {
		fmt.Println("error updating task:", err)
		osExit1()
	}

	if err := SaveTasks(taskFile, updatedTasks); err != nil {
		fmt.Println("error saving tasks:", err)
		osExit1()
	}

	fmt.Printf("Task %d updated successfully\n", id)
}

// handleDelete xử lý command: delete <id>
func handleDelete(tasks []Task, args []string) {
	if len(args) != 3 {
		fmt.Println(`usage: go run . delete <id>`)
		osExit1()
	}

	id, err := ParseID(args[2])
	if err != nil {
		fmt.Println("error:", err)
		osExit1()
	}

	updatedTasks, err := DeleteTask(tasks, id)
	if err != nil {
		fmt.Println("error deleting task:", err)
		osExit1()
	}

	if err := SaveTasks(taskFile, updatedTasks); err != nil {
		fmt.Println("error saving tasks:", err)
		osExit1()
	}

	fmt.Printf("Task %d deleted successfully\n", id)
}

// handleMarkInProgress xử lý command: mark-in-progress <id>
func handleMarkInProgress(tasks []Task, args []string) {
	if len(args) != 3 {
		fmt.Println(`usage: go run . mark-in-progress <id>`)
		osExit1()
	}

	id, err := ParseID(args[2])
	if err != nil {
		fmt.Println("error:", err)
		osExit1()
	}

	updatedTasks, err := MarkTaskStatus(tasks, id, StatusInProgress)
	if err != nil {
		fmt.Println("error updating task status:", err)
		osExit1()
	}

	if err := SaveTasks(taskFile, updatedTasks); err != nil {
		fmt.Println("error saving tasks:", err)
		osExit1()
	}

	fmt.Printf("Task %d marked as in-progress\n", id)
}

// handleMarkDone xử lý command: mark-done <id>
func handleMarkDone(tasks []Task, args []string) {
	if len(args) != 3 {
		fmt.Println(`usage: go run . mark-done <id>`)
		osExit1()
	}

	id, err := ParseID(args[2])
	if err != nil {
		fmt.Println("error:", err)
		osExit1()
	}

	updatedTasks, err := MarkTaskStatus(tasks, id, StatusDone)
	if err != nil {
		fmt.Println("error updating task status:", err)
		osExit1()
	}

	if err := SaveTasks(taskFile, updatedTasks); err != nil {
		fmt.Println("error saving tasks:", err)
		osExit1()
	}

	fmt.Printf("Task %d marked as done\n", id)
}

// handleList xử lý:
//
//	list
//	list done
//	list todo
//	list in-progress
func handleList(tasks []Task, args []string) {
	if len(args) > 3 {
		fmt.Println(`usage: go run . list [done|todo|in-progress]`)
		osExit1()
	}

	status := ""
	if len(args) == 3 {
		status = args[2]
	}

	filteredTasks, err := ListTasks(tasks, status)
	if err != nil {
		fmt.Println("error listing tasks:", err)
		osExit1()
	}

	PrintTasks(filteredTasks)
}

func osExit1() {
	panic(exitCode1{})
}

type exitCode1 struct{}
