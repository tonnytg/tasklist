package api_test

import (
	"bytes"
	"github.com/tonnytg/tasklist/pkg/api"
	"net/http"
	"os"
	"testing"
)

func TestGetTaskHandler(t *testing.T) {
	// set variable PORT to 9001 to run test
	os.Setenv("PORT", ":9001")
	go api.Start()

	// create a new task with POST call endpoint /api/task/add
	bj := bytes.NewReader([]byte(`{"name":"test","description":"test","status":"done"}`))
	resp, err := http.Post("http://localhost:9001/api/task/add", "application/json", bj)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Error("Error to create task, status retorned:", resp.StatusCode)
	}
}
