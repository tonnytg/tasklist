package api

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/tonnytg/tasklist/entities"
	"github.com/tonnytg/tasklist/internal/database"
	"log"
	"net/http"
	"os"
)

func Start() {

	http.HandleFunc("/api/tasks", ListHandler)
	http.HandleFunc("/api/tasks/add", CreateHandler)
	http.HandleFunc("/api/tasks/update/{id}", UpdateHandler)
	http.HandleFunc("/api/tasks/delete/{id}", DeleteHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":9000"
	}
	log.Println("Start API - port:", port)
	if err := http.ListenAndServe(port, nil); err == nil {
		panic("Api server can't running")
	}
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tasks := database.ListTask()
		for i,v := range tasks {
			fmt.Fprintf(w, "Tasks %d: %s \t description: %s \n", i, v.Name, v.Description)
		}
	}
	return
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		//createAt := time.Now().Format("2006-01-02")
		uuid, _ := uuid.NewRandom()
		name := fmt.Sprintf("TestTask-%s", uuid)
		task := entities.Task{Name: name, Description: "Task criada para teste", Status: true}

		// Save at database
		database.CreateTask(task.Name, task.Description, task.Status)

		tasks := database.ListTask()
		for i,v := range tasks {
			fmt.Fprintf(w, "Tasks %d: %s \t description: %s \n", i, v.Name, v.Description)
		}
	}
	return
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		tasks := database.ListTask()
		for i,v := range tasks {
			fmt.Fprintf(w, "Tasks %d: %s \t description: %s \n", i, v.Name, v.Description)
		}
	}
	return
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		tasks := database.ListTask()
		for i,v := range tasks {
			fmt.Fprintf(w, "Tasks %d: %s \t description: %s \n", i, v.Name, v.Description)
		}
	}
	return
}
