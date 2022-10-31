package database

import (
	"database/sql"
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

type TaskDb struct {
	db *sql.DB
}

func NewTaskDb(db *sql.DB) *TaskDb {
	return &TaskDb{db: db}
}

func (t *TaskDb) Get(ID uint16) (entities.TaskInterface, error) {
	var task entities.Task
	selectQuery := `select id, name, description, status from tasks where id = ?`
	stmt, err := t.db.Prepare(selectQuery)
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(ID).Scan(&task.ID, &task.Name, &task.Description, &task.Status)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (t *TaskDb) Save(task entities.TaskInterface) (entities.TaskInterface, error) {
	var rows int
	t.db.QueryRow("SELECT id FROM tasks WHERE id = ?", task.GetID()).Scan(&rows)
	if rows == 0 {
		_, err := t.create(task)
		if err != nil {
			return nil, err
		}
	}
	_, err := t.update(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t *TaskDb) create(task entities.TaskInterface) (entities.TaskInterface, error) {
	insertQuery := `insert into tasks (id, name, description, status) values (?, ?, ?, ?)`
	stmt, err := t.db.Prepare(insertQuery)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(
		task.GetID(),
		task.GetName(),
		task.GetDescription(),
		task.GetStatus(),
	)
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t *TaskDb) update(task entities.TaskInterface) (entities.TaskInterface, error) {
	updateQuery := `update tasks set name = ?, description = ?, status = ? where id = ?`
	stmt, err := t.db.Prepare(updateQuery)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(
		task.GetName(),
		task.GetDescription(),
		task.GetStatus(),
		task.GetID(),
	)
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return task, nil
}
