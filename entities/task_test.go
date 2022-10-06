package entities

import (
	"testing"
	"time"
)

func TestTask(t *testing.T) {
	task := NewTask()

	if task.ID != 0 {
		t.Errorf("Task ID is not 0")
	}
	if task.Hash != "" {
		t.Errorf("Task Hash is not empty")
	}
	if task.Name != "" {
		t.Errorf("Task Name is not empty")
	}
	if task.Description != "" {
		t.Errorf("Task Description is not empty")
	}
	if task.Status != 0 {
		t.Errorf("Task Status is not 0")
	}
	if task.CreatedAt != (time.Time{}) {
		t.Errorf("Task CreatedAt is not empty")
	}
	if task.UpdatedAt != (time.Time{}) {
		t.Errorf("Task UpdatedAt is not empty")
	}
	if task.ConvertTaskStatus() == "" {
		t.Errorf("Task ConvertTaskStatus can not empty")
	}
	if task.ConvertTaskStatus() != "Backlog" {
		t.Errorf("Task ConvertTaskStatus is not Backlog")
	}
}