package webserver

import (
	"html/template"
	"log"
	"net/http"
)

type Pessoa struct {
	Name string
	Age int32
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseGlob("./internal/webserver/templates/*.tmpl"))

	names := []Pessoa{Pessoa{"Antonio", 33}}

	err := tmpl.ExecuteTemplate(w, "Index", names)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}
