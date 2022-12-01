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
	http.HandleFunc("/api/task/add", Create)
	http.HandleFunc("/api/task/update", Update)
}

type TaskStruct struct {
	Full string        `json:"full"`
	Task entities.Task `json:"task"`
}

type Answer struct {
	Task   []entities.Task `json:"task"`
	Answer string          `json:"answer"`
	Status int             `json:"statusCode"`
}

func Get(w http.ResponseWriter, r *http.Request) {

	var a Answer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		hash := r.URL.Query().Get("hash")
		if hash == "" {
			a.Answer = "Hash is required"
			a.Status = http.StatusBadRequest
			jsonResp, _ := json.Marshal(a)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonResp)
			return
		}

		Con, err := database.NewTaskDb()
		if err != nil {
			errorMsg := fmt.Sprintf("cannot connect database:", err)
			a.Answer = errorMsg
			a.Status = http.StatusInternalServerError
			fmt.Fprintf(w, "status error:%d", a.Status)
			return
		}

		TaskService := entities.NewTaskService(Con)
		t, err := TaskService.Get(hash)
		if err != nil {
			log.Println("cannot get task:", err)
			a.Answer = err.Error()
			a.Status = http.StatusBadRequest
			fmt.Fprintf(w, "status: %d", a.Status)
			return
		}

		jsonResp, err := json.Marshal(t)
		if err != nil {
			log.Printf("error happened in json marshal list tasks; error: %s", err)
			a.Answer = err.Error()
			a.Status = http.StatusInternalServerError
			fmt.Fprintf(w, "status: %d", a.Status)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
		return
	}
	a.Status = http.StatusBadRequest
	a.Answer = "Bad request"
	fmt.Fprintf(w, "status: %d", a.Status)
	return
}

func Create(w http.ResponseWriter, r *http.Request) {

	var a Answer
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

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
		t, err := TaskService.Create(task.Name, task.Description, task.Body, task.Status)
		if err != nil {
			log.Println("cannot create task:", err)
			a.Answer = err.Error()
			a.Status = http.StatusBadRequest
			jsonResp, _ := json.Marshal(a)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonResp)
			return
		}
		jsonResp, err := json.Marshal(t)
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResp)
		fmt.Println(string(jsonResp))
		return
	}
	a.Status = http.StatusBadRequest
	a.Answer = "Bad request"
	fmt.Fprintf(w, "status: %d", a.Status)
	return
}

// TODO: update task by ID generate a new hash in persistence
// this will create a new taks not update, needs refactor

func Update(w http.ResponseWriter, r *http.Request) {

	var TaskTemp entities.Task
	reader, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(reader, &TaskTemp)

	Con, err := database.NewTaskDb()
	if err != nil {
		log.Println("cannot connect database:", err)
	}
	TaskService := entities.NewTaskService(Con)
	TaskService.Update(TaskTemp.Hash, TaskTemp.Name, TaskTemp.Description, TaskTemp.Body, TaskTemp.Status)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte("update task"))
	return
}
