package api

import (
	"encoding/json"
	"fmt"
	"github.com/tonnytg/tasklist/entities"
	"github.com/tonnytg/tasklist/internal/database"
	"io"
	"log"
	"net/http"
	"strconv"
)

func LoadHandlers() {
	http.HandleFunc("/api/task", GetTask)
	http.HandleFunc("/api/tasks", ListTasks)
	http.HandleFunc("/api/task/add", CreateTask)
	http.HandleFunc("/api/task/update", UpdateTaskByID)
	http.HandleFunc("/api/task/update_hash", UpdateTaskByHash)
	http.HandleFunc("/api/task/delete", DeleteTaks)
	http.HandleFunc("/api/tasks/delete/all", DeleteAllTasks)
}

func ListTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tasks, err := database.ListTask()
		if err != nil {
			log.Println(err)
		}

		TasksStruct := []struct {
			Full string        `json:"full"`
			Task entities.Task `json:"task"`
		}{}

		for _, v := range tasks {

			// fTask contains information from task to convert to json
			fTask := struct {
				Full string        `json:"full"`
				Task entities.Task `json:"task"`
			}{}
			fTask.Full = fmt.Sprintf("TaskID %d - %s - Status: %s",
				v.ID, v.Name, v.Status)
			fTask.Task = v

			TasksStruct = append(TasksStruct, fTask)
		}
		jsonResp, err := json.Marshal(TasksStruct)
		if err != nil {
			log.Printf("Error happened in JSON marshal list tasks. Err: %s", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
	}
	return
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		// convert url query to variable
		id := r.URL.Query().Get("id")

		// convert string to uint64
		tmpSearchID, err := strconv.ParseUint(id, 16, 16)
		// convert uint64 to uint16
		searchID := uint16(tmpSearchID)

		t, err := database.GetTaskByID(uint16(searchID))
		if err != nil {
			log.Println(err)
		}

		// fTask contains information from task to convert to json
		fTask := struct {
			Full string        `json:"full"`
			Task entities.Task `json:"task"`
		}{}
		fTask.Full = fmt.Sprintf("TaskID %d - %s - Status: %s",
			t.ID, t.Name, t.Status)
		fTask.Task = *t

		jsonResp, err := json.Marshal(fTask)
		if err != nil {
			log.Printf("Error happened in JSON marshal list tasks. Err: %s", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
	}
	return
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		// Example time format without millisecond if you needed
		// createAt := time.Now().Format("2006-01-02")

		reader, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}

		var t *entities.Task
		json.Unmarshal(reader, &t)
		task := entities.NewTask()
		task.SetName(t.Name)
		task.SetDescription(t.Description)
		task.SetStatus(t.Status)

		// Save at database
		t, err = database.CreateTask(*task)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			resp := make(map[string]string)
			resp["message"] = "Status Failed"
			jsonRespStatus, _ := json.Marshal(resp)
			w.Write(jsonRespStatus)
			return
		}

		// receive task from database
		jsonRespTask, err := json.Marshal(t)
		if err != nil {
			log.Printf("Error happened in JSON marshal. Err: %s", err)
		}

		resp := make(map[string]string)
		resp["message"] = "Status OK"
		jsonRespStatus, _ := json.Marshal(resp)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonRespTask)
		w.Write(jsonRespStatus)
	}
	return
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
