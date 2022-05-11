package database

import (
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

func CreateTask(name string, description string, status bool) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("func create failed to connect database")
	}
	db.Create(&entities.Task{
		Name:        name,
		Description: description,
		Status:      status,
	})
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
