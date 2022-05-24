package api

import (
	"encoding/json"
	"fmt"
	"github.com/tonnytg/tasklist/entities"
	"github.com/tonnytg/tasklist/internal/database"
	"io"
	"log"
	"net/http"
	"os"
)

func Start() {

	http.HandleFunc("/api/tasks", ListHandler)
	http.HandleFunc("/api/tasks/add", CreateHandler)
	http.HandleFunc("/api/tasks/update", UpdateHandler)
	http.HandleFunc("/api/tasks/delete/{id}", DeleteHandler)
	http.HandleFunc("/api/tasks/delete/all", DeleteAllHandler)

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
		for _, v := range tasks {

			// fTask contains information from task to convert to json
			fTask := struct {
				Full    string `json:"full"`
				Task   entities.Task `json:"task"`
			}{}
			fTask.Full = fmt.Sprintf("TaskID %d - %s - Status: %s",
				v.ID, v.Name, v.ConvertTaskStatus())
			fTask.Task = v

			jsonResp, err := json.Marshal(fTask)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal list tasks. Err: %s", err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResp)
		}
	}
	return
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		// Example time format without millisecond if you needed
		// createAt := time.Now().Format("2006-01-02")

		reader, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		var t entities.Task
		json.Unmarshal(reader, &t)
		task := entities.Task{Name: t.Name, Description: t.Description, Status: t.Status}

		// Save at database
		t, err = database.CreateTask(task.Name, task.Description, task.Status)
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
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
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

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		database.UpdateTask(10, "codigo", "developer")
		w.WriteHeader(200)
		w.Write([]byte("test"))
	}
	return
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		tasks := database.ListTask()
		for i, v := range tasks {
			fmt.Fprintf(w, "Tasks %d: %s \t description: %s \n", i, v.Name, v.Description)
		}
	}
	return
}

func DeleteAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		database.DeleteAllTasks()

		fmt.Fprintf(w, "Done")
	}
	return
}
