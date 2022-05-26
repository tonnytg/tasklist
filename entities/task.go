package entities

import (
	"time"
)

// status code
const (
	StatusBacklog  = 0 // Task in backlog
	StatusDoing    = 2 // Task moved from backlog to doing
	StatusDone     = 1 // Task finished
	StatusCanceled = 3 // Task canceled
)

type Task struct {
	ID          uint16      `json:"id"`
	Hash        string    `json:"hash"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ConvertTaskStatus translate status from code to string
func (t Task) ConvertTaskStatus() string {
	switch t.Status {
	case 0:
		return "Backlog"
	case 1:
		return "Doing"
	case 2:
		return "Done"
	case 3:
		return "Canceled"
	}
	return ""
}
