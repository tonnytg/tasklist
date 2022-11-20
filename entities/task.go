package entities

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

type Task struct {
	ID          uint16    `json:"id"`
	Hash        string    `json:"hash"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	BodyID      uint16    `json:"body_id"`
	Body        Body      `json:"body" gorm:"foreignKey:BodyID" gorm:"references:TaskHash"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Body struct {
	gorm.Model
	TaskHash string `json:"task_hash"`
	Content  string `json:"content"`
}

type TaskInterface interface {
	GetID() uint16
	GetIdByString() string
	GetHash() string
	GetName() string
	GetDescription() string
	GetBody() Body
	GetStatus() string
	SetHash(hash string) error
	SetName(name string) error
	SetDescription(description string) error
	SetBody(body Body) error
	SetStatus(status string) error
}

type TaskReader interface {
	Get(hash string) (TaskInterface, error)
}

type TaskWriter interface {
	Save(task *Task) (TaskInterface, error)
}

type TaskPersistenceInterface interface {
	TaskReader
	TaskWriter
}

func NewTask(name string, description string, body Body, status string) (*Task, error) {
	var task Task
	err := task.SetName(name)
	if err != nil {
		return nil, err
	}
	err = task.SetDescription(description)
	if err != nil {
		return nil, err
	}
	body.TaskHash = task.Hash
	err = task.SetBody(body)
	if err != nil {
		return nil, err
	}
	err = task.SetStatus(status)
	if err != nil {
		return nil, err
	}

	task.SetHash(uuid.New().String())
	return &task, nil
}

func (t *Task) GetID() uint16 {
	return t.ID
}

func (t *Task) GetIdByString() string {
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

func (t *Task) GetBody() Body {
	return t.Body
}

func (t *Task) GetStatus() string {
	return t.Status
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
		t.Description = description
		return nil
	}
	return errors.New("description has invalid characters or less than 50 characters")
}

func (t *Task) SetBody(body Body) error {
	t.Body = body
	return nil
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
		return nil
	}
	return errors.New("status is invalid")
}
