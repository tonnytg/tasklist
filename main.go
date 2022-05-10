package main

import (
	"fmt"
	"github.com/tonnytg/tasklist/entities"
	"github.com/tonnytg/tasklist/internal/webserver"
	"time"
)

func main() {
	fmt.Println("Tasklist with Go")

	createAt := time.Now().Format("2006-01-02")
	task := entities.Task{1, "2022", "Teste1", "Task criada para teste", true}

	fmt.Printf("---\nTask:\t\t%d\nName:\t\t%s\nCriada em:\t%s\n", task.ID, task.Name, createAt)

	// Start Webserver listening by default :8080
	webserver.Start()
}
