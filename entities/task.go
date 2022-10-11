package entities

import (
	"github.com/google/uuid"
	"strconv"
	"time"
)

const (
	BACKLOG  = "backlog"
	DOING    = "doing"
	DONE     = "done"
	CANCELED = "canceled"
)

type TaskInterface interface {
	GetID() uint16
	GetIDString() string
	GetHash() string
	SetName(name string)
	SetDescription(description string)
	SetStatus(status string)
}

// TaskServiceInterface orquest the task actions orders
type TaskServiceInterface interface {
	Create(name string, description string, status string) (Task, error)
	Get(ID uint16) Task
	Update(ID uint16, name string, description string) error
	Delete(ID uint16) error
}

type TaskReader interface {
	Get(ID uint16) (TaskInterface, error)
}

type TaskWriter interface {
	Save(task TaskInterface) (TaskInterface, error)
}

type TaskPersistenceInterface interface {
	TaskReader
	TaskWriter
}

type Task struct {
	ID          uint16    `json:"id"`
	Hash        string    `json:"hash"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewTask() *Task {
	t := Task{}
	t.Hash = uuid.NewString()
	t.Status = BACKLOG
	return &t
}

func (t *Task) GetID() uint16 {
	return t.ID
}

func (t *Task) GetIDString() string {
	i := strconv.Itoa(int(t.ID))
	return i
}

func (t *Task) GetHash() string {
	return t.Hash
}

func (t *Task) SetName(name string) {
	t.Name = name
}

func (t *Task) SetDescription(description string) {
	t.Description = description
}

func (t *Task) SetStatus(status string) {
	switch status {
	case DOING:
		t.Status = DOING
	case DONE:
		t.Status = DONE
	case CANCELED:
		t.Status = CANCELED
	default:
		t.Status = BACKLOG
	}
}
