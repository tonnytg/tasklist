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

func init() {

	db, err := gorm.Open(sqlite.Open(fullPath), &gorm.Config{})
	if err != nil {
		panic("func init failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&entities.Task{})
}

func CreateTask(task entities.Task) (*entities.Task, error) {

	db, err := gorm.Open(sqlite.Open(fullPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	tx := db.Create(&task)
	return &task, tx.Error
}

func GetTask(ID uint16) (*entities.Task, error) {
	db, err := gorm.Open(sqlite.Open(fullPath), &gorm.Config{})
	if err != nil {
		panic("func create failed to connect database")
	}

	tempTask := entities.Task{}
	db.First(&tempTask, ID)
	if tempTask.ID == 0 {
		return nil, errors.New("task not found")
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

func UpdateTask(ID uint16, name string, description string) (*entities.Task, error) {
	db, err := gorm.Open(sqlite.Open(fullPath), &gorm.Config{})
	if err != nil {
		panic("func create failed to connect database")
	}

	task, err := GetTask(ID)
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
