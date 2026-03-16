package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestLoadTasks_FileNotExist(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "tasks.json")

	tasks, err := LoadTasks(filename)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(tasks) != 0 {
		t.Errorf("expected empty task list, got %d tasks", len(tasks))
	}
}

func TestLoadTasks_EmptyFile(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "tasks.json")

	if err := os.WriteFile(filename, []byte(""), 0644); err != nil {
		t.Fatalf("failed to create empty file: %v", err)
	}

	tasks, err := LoadTasks(filename)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(tasks) != 0 {
		t.Errorf("expected empty task list, got %d tasks", len(tasks))
	}
}

func TestLoadTasks_InvalidJSON(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "tasks.json")

	if err := os.WriteFile(filename, []byte(`{"invalid":`), 0644); err != nil {
		t.Fatalf("failed to write invalid json: %v", err)
	}

	_, err := LoadTasks(filename)
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
}

func TestSaveTasksAndLoadTasks(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "tasks.json")

	now := time.Now()

	originalTasks := []Task{
		{
			ID:          1,
			Description: "Learn Go",
			Status:      StatusTodo,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			ID:          2,
			Description: "Write tests",
			Status:      StatusDone,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}

	if err := SaveTasks(filename, originalTasks); err != nil {
		t.Fatalf("expected no error saving tasks, got %v", err)
	}

	loadedTasks, err := LoadTasks(filename)
	if err != nil {
		t.Fatalf("expected no error loading tasks, got %v", err)
	}

	if len(loadedTasks) != len(originalTasks) {
		t.Fatalf("expected %d tasks, got %d", len(originalTasks), len(loadedTasks))
	}

	for i := range originalTasks {
		if loadedTasks[i].ID != originalTasks[i].ID {
			t.Errorf("task %d: expected ID %d, got %d", i, originalTasks[i].ID, loadedTasks[i].ID)
		}

		if loadedTasks[i].Description != originalTasks[i].Description {
			t.Errorf(
				"task %d: expected description %q, got %q",
				i,
				originalTasks[i].Description,
				loadedTasks[i].Description,
			)
		}

		if loadedTasks[i].Status != originalTasks[i].Status {
			t.Errorf("task %d: expected status %q, got %q", i, originalTasks[i].Status, loadedTasks[i].Status)
		}

		if loadedTasks[i].CreatedAt.IsZero() {
			t.Errorf("task %d: expected CreatedAt to be set", i)
		}

		if loadedTasks[i].UpdatedAt.IsZero() {
			t.Errorf("task %d: expected UpdatedAt to be set", i)
		}
	}
}
