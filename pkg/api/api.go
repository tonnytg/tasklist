package api

import (
	"fmt"
	"github.com/tonnytg/tasklist/internal/database"
	"log"
	"net/http"
	"os"
)

func Start() {

	http.HandleFunc("/", ListHandle)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":9000"
	}
	log.Println("Start API - port:", port)
	if err := http.ListenAndServe(port, nil); err == nil {
		panic("Api server can't running")
	}
}

func ListHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tasks := database.ListTask()
		for i,v := range tasks {
			fmt.Fprintf(w, "Tasks %d: %s \t description: %s \n", i, v.Name, v.Description)
		}
	}
	return
}
