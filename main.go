package main

import (
	"fmt"
	"github.com/tonnytg/tasklist/entities"
	"github.com/tonnytg/tasklist/internal/database"
	"github.com/tonnytg/tasklist/pkg/api"
)

func main() {
	fmt.Println("Tasklist with Go")

	//createAt := time.Now().Format("2006-01-02")
	task := entities.Task{Name: "Teste1", Description: "Task criada para teste", Status: true}

	// Save at database
	database.CreateTask(task.Name, task.Description, task.Status)

	task2 := database.GetTask(1) // find product with integer primary key
	fmt.Println("Task geted:", task2)

	// Start API to listening
	api.Start()
}
