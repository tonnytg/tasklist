package webserver

import (
	"net/http"
)

func routes() {
	http.HandleFunc("/", indexHandler)
}
