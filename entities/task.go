package entities

import (
	"github.com/google/uuid"
	"time"
)

const (
	StatusBacklog  = 0 // Task in backlog
	StatusDoing    = 2 // Task moved from backlog to doing
	StatusDone     = 1 // Task finished
	StatusCanceled = 3 // Task canceled

	BACKLOG  = "backlog"
	DOING    = "doing"
	DONE     = "done"
	CANCELED = "canceled"
	EMPTY    = "empty"
)

type Task struct {
	ID          uint16    `json:"id"`
	Hash        string    `json:"hash"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskInterface interface {
	GetTask() *Task
	GetID() uint16
	GetStatus() string
	GetHash() string
	GetName() string
}

func NewTask() *Task {
	t := Task{}
	t.Hash = uuid.NewString()
	t.Status = BACKLOG
	return &t
}

func (t *Task) GetTask() *Task {
	return t
}

func (t *Task) GetStatus() string {
	return t.Status
}

func (t *Task) GetID() uint16 {
	return t.ID
}

func (t *Task) GetHash() string {
	return t.Hash
}

func (t *Task) GetName() string {
	return t.Name
}
