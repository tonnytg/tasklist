package database

import (
	"errors"
	"fmt"
	"github.com/tonnytg/tasklist/entities"
	"github.com/tonnytg/tasklist/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	p        = config.Reader()
	fullPath = p.Database.Path + "/" + p.Database.Name
)

type TaskDb struct {
	db *gorm.DB
}

func init() {

	Con, err := NewTaskDb()
	if err != nil {
		panic("failed to connect database in init func" + err.Error())
	}

	// Migrate the schema
	Con.db.AutoMigrate(&entities.Task{})
}

func NewTaskDb() (*TaskDb, error) {

	var db *gorm.DB
	db, err := gorm.Open(sqlite.Open(fullPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &TaskDb{db: db}, err
}

func (td *TaskDb) Get(hash string) (entities.TaskInterface, error) {

	tempTask := entities.Task{}
	td.db.Where("hash = ?", hash).First(&tempTask)
	if tempTask.ID == 0 {
		return nil, errors.New("task by hash not found")
	}
	return &tempTask, nil
}

func (td *TaskDb) create(task *entities.Task) (entities.TaskInterface, error) {

	result := td.db.Create(&task)
	if result.Error != nil {
		errorMsg := fmt.Sprintf("failed to create task: %s", result.Error.Error())
		return nil, errors.New(errorMsg)
	}

	return task, nil
}

func (td *TaskDb) update(task *entities.Task) (entities.TaskInterface, error) {

	var tempTask entities.Task

	result := td.db.Where("hash = ?", task.Hash).First(&tempTask)
	if result.Error != nil {
		return nil, result.Error
	}

	tempTask.SetName(task.GetName())
	tempTask.SetDescription(task.GetDescription())
	tempTask.SetBody(task.GetBody())
	tempTask.SetStatus(task.GetStatus())

	result = td.db.Save(&tempTask)
	if result.Error != nil {
		errorMsg := fmt.Sprintf("failed to update task: %s", result.Error.Error())
		return nil, errors.New(errorMsg)
	}

	return task, nil
}

func (td *TaskDb) Save(task *entities.Task) (entities.TaskInterface, error) {

	var tempTask entities.Task
	result := td.db.Where("hash = ?", task.GetHash()).First(&tempTask)
	if result.Error != nil {
		_, err := td.create(task)
		if err != nil {
			return nil, err
		}
		return task, nil
	}

	_, err := td.update(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}
