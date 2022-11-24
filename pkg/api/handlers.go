package api

import (
	"encoding/json"
	"fmt"
	"github.com/tonnytg/tasklist/entities"
	"github.com/tonnytg/tasklist/internal/database"
	"io"
	"log"
	"net/http"
)

func LoadHandlers() {
	http.HandleFunc("/api/task", Get)
	http.HandleFunc("/api/tasks", ListTasks)
	http.HandleFunc("/api/task/add", Create)
	http.HandleFunc("/api/task/update", UpdateTaskByID)
	http.HandleFunc("/api/task/update_hash", UpdateTaskByHash)
	http.HandleFunc("/api/task/delete", DeleteTaks)
	http.HandleFunc("/api/tasks/delete/all", DeleteAllTasks)
}

type TaskStruct struct {
	Full string        `json:"full"`
	Task entities.Task `json:"task"`
}

func ListTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tasks, err := database.ListTask()
		if err != nil {
			log.Println(err)
		}

		var tasksStruct []TaskStruct

		for _, v := range tasks {

			// fTask contains information from task to convert to json
			var t TaskStruct
			t.Full = fmt.Sprintf("TaskID %d - %s - Status: %s",
				v.ID, v.Name, v.Status)
			t.Task = v

			tasksStruct = append(tasksStruct, t)
		}
		jsonResp, err := json.Marshal(tasksStruct)
		if err != nil {
			log.Printf("Error happened in JSON marshal list tasks. Err: %s", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
	}
	return
}

func Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		hash := r.URL.Query().Get("hash")

		Con, err := database.NewTaskDb()
		if err != nil {
			log.Println("cannot connect database:", err)
		}

		TaskService := entities.NewTaskService(Con)
		t, _ := TaskService.Get(hash)
		fmt.Println("t:", t)
		jsonResp, err := json.Marshal(t)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
		fmt.Println(string(jsonResp))
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		reader, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}

		task := entities.Task{}
		json.Unmarshal(reader, &task)

		Con, err := database.NewTaskDb()
		if err != nil {
			log.Println("cannot connect database:", err)
		}

		TaskService := entities.NewTaskService(Con)
		t, _ := TaskService.Create(task.Name, task.Description, task.Body, task.Status)
		jsonResp, err := json.Marshal(t)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResp)
		fmt.Println(string(jsonResp))
	}
}

func UpdateTaskByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		reader, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		var t *entities.Task
		json.Unmarshal(reader, &t)

		database.UpdateTaskByID(t.ID, t.Name, t.Description)
		w.WriteHeader(200)
		w.Write([]byte("Success"))
	}
	return
}

func UpdateTaskByHash(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		reader, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		var t *entities.Task
		json.Unmarshal(reader, &t)

		database.UpdateTaskByHash(t.Hash, t.Name, t.Description)
		w.WriteHeader(200)
		w.Write([]byte("Success"))
	}
	return
}

func DeleteTaks(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {

		reader, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}

		var t entities.Task
		json.Unmarshal(reader, &t)

		database.DeleteTask(t)

		fmt.Fprintf(w, "Done")
	}
	return
}

func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		database.DeleteAllTasks()

		fmt.Fprintf(w, "Done")
	}
	return
}
