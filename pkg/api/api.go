package api

import (
	"log"
	"net/http"
	"os"
)

func Start() {

	LoadHandlers()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":9000"
	}
	log.Println("Start API - port:", port)
	if err := http.ListenAndServe(port, nil); err == nil {
		panic("Api server can't running")
	}
}
