package binary

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tonnytg/tasklist/entities"
)

func WebGet() {
	resp, err := http.Get("http://localhost:9000/api/tasks")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	type Tasks []struct {
		Full string        `json:"full"`
		Task entities.Task `json:"task"`
	}

	ts := Tasks{}
	json.NewDecoder(resp.Body).Decode(&ts)
	var count int
	for i, v := range ts {
		fmt.Printf("[%d] - %s:\t %s\n", i, v.Task.Name, v.Task.Description)
		count++
	}
	if count == 0 {
		fmt.Println("No tasks found")
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
	fmt.Println("List Tasks:")
	WebGet()
}

func Update() {
	fmt.Println("Updated Task")
}

func Delete() {
	fmt.Println("Deleted Task")
}
