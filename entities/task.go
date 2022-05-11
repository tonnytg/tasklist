package entities

import (
	"time"
)

type Task struct {
	ID          uint
	Name        string
	Description string
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
