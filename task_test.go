package main

import "testing"

func TestNewTask(t *testing.T) {
	task, err := NewTask(1, "Learn Go")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if task.ID != 1 {
		t.Errorf("expected ID 1, got %d", task.ID)
	}

	if task.Description != "Learn Go" {
		t.Errorf("expected description %q, got %q", "Learn Go", task.Description)
	}

	if task.Status != StatusTodo {
		t.Errorf("expected status %q, got %q", StatusTodo, task.Status)
	}

	if task.CreatedAt.IsZero() {
		t.Error("expected CreatedAt to be set")
	}

	if task.UpdatedAt.IsZero() {
		t.Error("expected UpdatedAt to be set")
	}
}

func TestNewTask_EmptyDescription(t *testing.T) {
	_, err := NewTask(1, "   ")
	if err == nil {
		t.Fatal("expected error for empty description, got nil")
	}

	if err.Error() != "description cannot be empty" {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestNextTaskID(t *testing.T) {
	tasks := []Task{
		{ID: 1},
		{ID: 3},
		{ID: 7},
	}

	got := NextTaskID(tasks)
	want := 8

	if got != want {
		t.Errorf("expected next ID %d, got %d", want, got)
	}
}

func TestAddTask(t *testing.T) {
	var tasks []Task

	updatedTasks, task, err := AddTask(tasks, "Build CLI")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(updatedTasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(updatedTasks))
	}

	if task.ID != 1 {
		t.Errorf("expected task ID 1, got %d", task.ID)
	}

	if updatedTasks[0].Description != "Build CLI" {
		t.Errorf("expected description %q, got %q", "Build CLI", updatedTasks[0].Description)
	}
}

func TestAddTask_EmptyDescription(t *testing.T) {
	var tasks []Task

	_, _, err := AddTask(tasks, "   ")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestUpdateTask(t *testing.T) {
	tasks := []Task{
		{
			ID:          1,
			Description: "Old description",
			Status:      StatusTodo,
		},
	}

	updatedTasks, err := UpdateTask(tasks, 1, "New description")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if updatedTasks[0].Description != "New description" {
		t.Errorf("expected updated description, got %q", updatedTasks[0].Description)
	}

	if updatedTasks[0].UpdatedAt.IsZero() {
		t.Error("expected UpdatedAt to be set")
	}
}

func TestUpdateTask_NotFound(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Task 1"},
	}

	_, err := UpdateTask(tasks, 99, "New")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "task not found" {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestUpdateTask_EmptyDescription(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Task 1"},
	}

	_, err := UpdateTask(tasks, 1, "   ")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "new description cannot be empty" {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestUpdateTask_SameDescription(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Task 1"},
	}

	_, err := UpdateTask(tasks, 1, "Task 1")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "new description is the same as the current one" {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestDeleteTask(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Task 1"},
		{ID: 2, Description: "Task 2"},
		{ID: 3, Description: "Task 3"},
	}

	updatedTasks, err := DeleteTask(tasks, 2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(updatedTasks) != 2 {
		t.Fatalf("expected 2 tasks after delete, got %d", len(updatedTasks))
	}

	if updatedTasks[0].ID != 1 || updatedTasks[1].ID != 3 {
		t.Errorf("unexpected task order after delete: %+v", updatedTasks)
	}
}

func TestDeleteTask_NotFound(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Task 1"},
	}

	_, err := DeleteTask(tasks, 99)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "task not found" {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestMarkTaskStatus(t *testing.T) {
	tasks := []Task{
		{
			ID:          1,
			Description: "Task 1",
			Status:      StatusTodo,
		},
	}

	updatedTasks, err := MarkTaskStatus(tasks, 1, StatusDone)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if updatedTasks[0].Status != StatusDone {
		t.Errorf("expected status %q, got %q", StatusDone, updatedTasks[0].Status)
	}

	if updatedTasks[0].UpdatedAt.IsZero() {
		t.Error("expected UpdatedAt to be set")
	}
}

func TestMarkTaskStatus_InvalidStatus(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Task 1", Status: StatusTodo},
	}

	_, err := MarkTaskStatus(tasks, 1, "finished")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "invalid status" {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestMarkTaskStatus_SameStatus(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Task 1", Status: StatusDone},
	}

	_, err := MarkTaskStatus(tasks, 1, StatusDone)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "task already has this status" {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestMarkTaskStatus_NotFound(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Task 1", Status: StatusTodo},
	}

	_, err := MarkTaskStatus(tasks, 99, StatusDone)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "task not found" {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestListTasks_All(t *testing.T) {
	tasks := []Task{
		{ID: 1, Status: StatusTodo},
		{ID: 2, Status: StatusDone},
	}

	got, err := ListTasks(tasks, "")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(got) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(got))
	}
}

func TestListTasks_Filtered(t *testing.T) {
	tasks := []Task{
		{ID: 1, Status: StatusTodo},
		{ID: 2, Status: StatusDone},
		{ID: 3, Status: StatusDone},
	}

	got, err := ListTasks(tasks, StatusDone)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(got) != 2 {
		t.Fatalf("expected 2 done tasks, got %d", len(got))
	}

	if got[0].ID != 2 || got[1].ID != 3 {
		t.Errorf("unexpected filtered tasks: %+v", got)
	}
}

func TestListTasks_InvalidFilter(t *testing.T) {
	tasks := []Task{
		{ID: 1, Status: StatusTodo},
	}

	_, err := ListTasks(tasks, "finished")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "invalid status filter" {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestIsValidStatus(t *testing.T) {
	validStatuses := []string{StatusTodo, StatusInProgress, StatusDone}
	for _, status := range validStatuses {
		if !IsValidStatus(status) {
			t.Errorf("expected status %q to be valid", status)
		}
	}

	invalidStatuses := []string{"", "finished", "pending"}
	for _, status := range invalidStatuses {
		if IsValidStatus(status) {
			t.Errorf("expected status %q to be invalid", status)
		}
	}
}
