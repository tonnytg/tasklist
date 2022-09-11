package binary

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func WebGet() {
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

func WebPost(name string, description string) {

	values := map[string]string{"Name": name, "Description": description}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:9000/api/tasks/add", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])
}

func Create(data string, description string) {
	fmt.Println("Created Task")
	WebPost(data, description)
}

func List() {
	fmt.Println("Listed Tasks")
	WebGet()
}

func Update() {
	fmt.Println("Updated Task")
}

func Delete() {
	fmt.Println("Deleted Task")
}
