package database

import (
	"github.com/google/uuid"
	"github.com/tonnytg/tasklist/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("func init failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&entities.Task{})
}

func CreateTask(name string, description string, status int) error {

	hash := uuid.NewString() // generate a unique hash to task

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("func create failed to connect database")
	}
	tx := db.Create(&entities.Task{
		Hash:        hash,
		Name:        name,
		Description: description,
		Status:      status,
	})
	return tx.Error
}

func GetTask(ID int32) entities.Task {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("func create failed to connect database")
	}

	tempTask := entities.Task{}
	db.First(&tempTask, ID)
	return tempTask
}

func ListTask() []entities.Task {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("func create failed to connect database")
	}

	var tempTask []entities.Task
	db.Find(&tempTask)
	return tempTask
}

func UpdateTask(ID int32, name string, description string) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("func create failed to connect database")
	}

	task := GetTask(ID)
	task.Name = name
	task.Description = description
	db.Save(&task)
}

func DeleteAllTasks() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("func create failed to connect database")
	}

	db.Where("status = ?", entities.StatusDone).Delete(&entities.Task{})
	db.Where("status = ?", entities.StatusCanceled).Delete(&entities.Task{})
	db.Where("status = ?", entities.StatusDoing).Delete(&entities.Task{})
	db.Where("status = ?", entities.StatusBacklog).Delete(&entities.Task{})
}
