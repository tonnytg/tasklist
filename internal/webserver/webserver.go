package webserver

import (
	"log"
	"net/http"
	"os"
)

func Start() {

	// to stay more clean all rotes are in separate package
	routes()

	// define port webserver if no port exported
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	// Load static files to easily way to make to us a show
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("."))))

	// log fatal if has a problem to open port :8080
	// if it ok with port, http load files inside http.Dir
	log.Fatal(http.ListenAndServe(port, nil))
}
