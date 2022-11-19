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

func (td TaskDb) update(task *entities.Task) (entities.TaskInterface, error) {

	result := td.db.Where("hash = ?", task.Hash).First(&task)
	if result.Error != nil {
		return nil, result.Error
	}

	result = td.db.Save(&task)
	if result.Error != nil {
		errorMsg := fmt.Sprintf("failed to update task: %s", result.Error.Error())
		return nil, errors.New(errorMsg)
	}

	return task, nil
}

func (td *TaskDb) Save(task *entities.Task) (entities.TaskInterface, error) {

	result := td.db.Where("hash = ?", task.GetHash()).First(&task)
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

func CreateTask(task entities.Task) (*entities.Task, error) {

	db, err := gorm.Open(sqlite.Open(fullPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	tx := db.Create(&task)
	return &task, tx.Error
}

// TODO: Remover conexões sem o serviceInterface
func GetTaskByID(ID uint16) (*entities.Task, error) {
	db, err := gorm.Open(sqlite.Open(fullPath), &gorm.Config{})
	if err != nil {
		panic("func create failed to connect database")
	}

	tempTask := entities.Task{}
	db.First(&tempTask, ID)
	if tempTask.ID == 0 {
		return nil, errors.New("task by id not found")
	}
	return &tempTask, nil
}

func GetTaskByHash(Hash string) (*entities.Task, error) {
	db, err := gorm.Open(sqlite.Open(fullPath), &gorm.Config{})
	if err != nil {
		panic("func create failed to connect database")
	}

	tempTask := entities.Task{}
	// db get by hash
	db.Where("hash = ?", Hash).First(&tempTask)
	if tempTask.ID == 0 {
		return nil, errors.New("task by hash not found")
	}
	return &tempTask, nil
}

func ListTask() ([]entities.Task, error) {
	db, err := gorm.Open(sqlite.Open(fullPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var tempTask []entities.Task
	db.Find(&tempTask)
	return tempTask, nil
}

func UpdateTaskByID(ID uint16, name string, description string) (*entities.Task, error) {
	db, err := gorm.Open(sqlite.Open(fullPath), &gorm.Config{})
	if err != nil {
		panic("func create failed to connect database")
	}

	task, err := GetTaskByID(ID)
	if err != nil {
		return nil, err
	}
	task.Name = name
	task.Description = description
	db.Save(&task)

	return task, nil
}

func UpdateTaskByHash(hash string, name string, description string) (*entities.Task, error) {
	db, err := gorm.Open(sqlite.Open(fullPath), &gorm.Config{})
	if err != nil {
		panic("func create failed to connect database")
	}

	task, err := GetTaskByHash(hash)
	if err != nil {
		return nil, err
	}
	task.Name = name
	task.Description = description
	db.Save(&task)

	return task, nil
}

func DeleteAllTasks() error {
	db, err := gorm.Open(sqlite.Open(fullPath), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %s", err)
	}

	db.Where("status = ?", entities.DONE).Delete(&entities.Task{})
	db.Where("status = ?", entities.CANCELED).Delete(&entities.Task{})
	db.Where("status = ?", entities.DOING).Delete(&entities.Task{})
	db.Where("status = ?", entities.BACKLOG).Delete(&entities.Task{})

	tasks := []entities.Task{}
	db.Find(&tasks)
	if len(tasks) != 0 {
		return errors.New("failed to delete all tasks")
	}
	return nil
}

func DeleteTask(task entities.Task) {
	db, err := gorm.Open(sqlite.Open(fullPath), &gorm.Config{})
	if err != nil {
		panic("func create failed to connect database")
	}

	db.Where("hash = ?", task.Hash).Delete(&entities.Task{})
}
