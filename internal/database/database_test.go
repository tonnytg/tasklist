package database_test

import (
	"database/sql"
	"github.com/tonnytg/tasklist/entities"
	"github.com/tonnytg/tasklist/internal/database"
	"log"
	"testing"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createTask(Db)
}

func createTable(db *sql.DB) {
	tableQuery := `CREATE table tasks (
		"id" string,
		"name" string,
		"description" string,
		"status" string
		);`

	stmt, err := db.Prepare(tableQuery)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createTask(db *sql.DB) {
	insertQuery := `INSERT INTO tasks (id, name, description, status) VALUES (?, ?, ?, ?);`
	stmt, err := db.Prepare(insertQuery)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec(1, "Task 1", "Description 1", entities.BACKLOG)
}

func TestTaskDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	taskDb := database.NewTaskDb(Db)
	task, err := taskDb.Get(1)
	if err != nil {
		t.Error("Error on get task", err)
	}
	if task.GetName() != "Task 1" {
		t.Error("Error on get task")
	}
	if task.GetStatus() != entities.BACKLOG {
		t.Error("Error on get task")
	}
}

func TestTaskDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	taskDb := database.NewTaskDb(Db)
	task := entities.NewTask()
	task.SetName("Task 1")
	task.SetStatus(entities.DOING)
	task.SetDescription("Description 1")
	result, err := taskDb.Save(task)
	if err != nil {
		t.Error("Error on save task", err)
	}
	if result.GetName() != "Task 1" {
		t.Error("Error on save task, name is not equal")
	}
	if result.GetStatus() != entities.DOING {
		t.Error("Error on save task, status is not equal")
	}

	task.SetStatus(entities.DONE)
	result, err = taskDb.Save(task)
	if err != nil {
		t.Error("Error on save task", err)
	}
	if result.GetStatus() != entities.DONE {
		t.Error("Error on save task")
	}
}
