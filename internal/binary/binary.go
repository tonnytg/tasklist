package binary

import (
	"bufio"
	"fmt"
	"net/http"
)

func WebRequest() {
	resp, err := http.Get("http://localhost:9000/api/tasks")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the HTTP response status.
	fmt.Println("Response status:", resp.Status)

	// Print the first 5 lines of the response body.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func Create() {
	fmt.Println("Created Task")
}

func List() {
	fmt.Println("Listed Tasks")
	WebRequest()
}

func Update() {
	fmt.Println("Updated Task")
}

func Delete() {
	fmt.Println("Deleted Task")
}
