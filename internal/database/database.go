package database

import (
	"database/sql"
	"fmt"
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

func CreateTask(name string, description string, status string) (entities.Task, error) {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("func create failed to connect database")
	}

	t := entities.NewTask()

	tx := db.Create(&t)
	return *t, tx.Error
}

func GetTask(ID uint16) entities.Task {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("func create failed to connect database")
	}

	tempTask := entities.Task{}
	db.First(&tempTask, ID)
	fmt.Println("Achei", tempTask)
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

func UpdateTask(ID uint16, name string, description string) {
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

	db.Where("status = ?", entities.DONE).Delete(&entities.Task{})
	db.Where("status = ?", entities.CANCELED).Delete(&entities.Task{})
	db.Where("status = ?", entities.DOING).Delete(&entities.Task{})
	db.Where("status = ?", entities.BACKLOG).Delete(&entities.Task{})
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
