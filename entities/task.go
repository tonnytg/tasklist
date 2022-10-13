package entities

import (
	"errors"
	"github.com/google/uuid"
	"regexp"
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
	GetName() string
	GetDescription() string
	GetStatus() string
	GetHash() string
	SetID(ID uint16) error
	SetHash(hash string) error
	SetName(name string) error
	SetDescription(description string) error
	SetStatus(status string) error
}

// TaskServiceInterface orquest the task actions orders
type TaskServiceInterface interface {
	Create(name string, description string, status string) (TaskInterface, error)
	Get(hash string) (TaskInterface, error)
	Update(hash string, name string, description string, status string) (TaskInterface, error)
}

type TaskReader interface {
	Get(hash string) (TaskInterface, error)
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

func (t *Task) GetName() string {
	return t.Name
}

func (t *Task) GetDescription() string {
	return t.Description
}

func (t *Task) GetStatus() string {
	return t.Status
}

func (t *Task) SetID(ID uint16) error {
	t.ID = ID
	return nil
}

func (t *Task) SetHash(hash string) error {
	hashConverted, err := uuid.Parse(hash)
	if err != nil {
		return err
	}
	if hash == "" {
		return errors.New("hash is empty")
	}
	t.Hash = hashConverted.String()
	return nil
}

func (t *Task) SetName(name string) error {
	regexName := regexp.MustCompile(`^[a-zA-Z0-9_ áàâãéèêíïóôõöúçñ]{1,30}$`)
	if regexName.MatchString(name) {
		t.Name = name
		return nil
	}
	return errors.New("name has invalid characters or less than 30 characters")
}

func (t *Task) SetDescription(description string) error {
	regexDesc := regexp.MustCompile(`^[a-zA-Z0-9_ áàâãéèêíïóôõöúçñ]{1,50}$`)
	if regexDesc.MatchString(description) {
		t.Name = description
		return nil
	}
	return errors.New("description has invalid characters or less than 50 characters")
}

func (t *Task) SetStatus(status string) error {
	switch status {
	case DOING:
		t.Status = DOING
		return nil
	case DONE:
		t.Status = DONE
		return nil
	case CANCELED:
		t.Status = CANCELED
		return nil
	default:
		t.Status = BACKLOG
	}
	return errors.New("status is invalid")
}
