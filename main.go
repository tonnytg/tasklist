package main

import (
	"fmt"
	"github.com/tonnytg/tasklist/pkg/api"
)

func main() {
	fmt.Println("Tasklist with Go")

	// Start API to listening
	api.Start()
}
