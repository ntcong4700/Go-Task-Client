package main

import (
	"fmt"
	"os"
)

const taskFile = "tasks.json"

func main() {
	if len(os.Args) < 2 {
		PrintUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	tasks, err := LoadTasks(taskFile)
	if err != nil {
		fmt.Println("error loading tasks:", err)
		os.Exit(1)
	}

	switch command {
	case "add":
		handleAdd(tasks, os.Args)
	case "update":
		handleUpdate(tasks, os.Args)
	case "delete":
		handleDelete(tasks, os.Args)
	case "mark-in-progress":
		handleMarkInProgress(tasks, os.Args)
	case "mark-done":
		handleMarkDone(tasks, os.Args)
	case "list":
		handleList(tasks, os.Args)
	default:
		fmt.Println("unknown command:", command)
		PrintUsage()
		os.Exit(1)
	}
}
