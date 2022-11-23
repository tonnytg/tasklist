package database_test

import (
	"github.com/tonnytg/tasklist/entities"
	"github.com/tonnytg/tasklist/internal/database"
	"testing"
)

func TestTaskDb_Get(t *testing.T) {

	// Create connection
	Db, err := database.NewTaskDb()
	if err != nil {
		t.Errorf("failed to connect database in init func" + err.Error())
	}

	// Create Task
	task, err := entities.NewTask("testName", "testDesc", entities.Body{Content: "body test"}, "backlog")
	if err != nil {
		t.Errorf("failed to create task: %s", err)
	}

	// Save Task
	_, err = Db.Save(task)
	if err != nil {
		t.Errorf("failed to save task in database:" + err.Error())
	}

	// Get Task
	_, err = Db.Get(task.GetHash())
	if err != nil {
		t.Errorf("failed to get task:" + err.Error())
	}
}
