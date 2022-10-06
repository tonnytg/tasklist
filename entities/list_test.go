package entities

import "testing"

func TestList(t *testing.T) {
	l := NewList()
	if l.ID != 0 {
		t.Errorf("Expected ID to be 0, got %d", l.ID)
	}
	if l.CreateAt != "" {
		t.Errorf("Expected CreateAt to be empty, got %s", l.CreateAt)
	}
	if l.Name != "" {
		t.Errorf("Expected Name to be empty, got %s", l.Name)
	}
	if len(l.Tasks) != 0 {
		t.Errorf("Expected Tasks to be empty, got %d", len(l.Tasks))
	}
}
