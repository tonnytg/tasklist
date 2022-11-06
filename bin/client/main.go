package main

import (
	"fmt"
	"os"
)

const ClientHelp string = `pass help as argument to get help.
./taskclient help`
const AddHelp string = `Help Add!`

func main() {
	fmt.Println("Client Side")

	if len(os.Args) < 2 {
		fmt.Println(ClientHelp)
		os.Exit(0)
	}

	os.Args = os.Args[1:]
	if os.Args[0] == "help" {
		fmt.Println(ClientHelp)
		os.Exit(0)
	}

	if os.Args[0] == "add" {
		if len(os.Args) < 2 {
			fmt.Println(AddHelp)
			os.Exit(1)
		}
		fmt.Println("Add")
		args := os.Args[1:]
		fmt.Println("lenArgs:", len(args))
		fmt.Println("Title:", args[0])
		if args[1] == "desc" || args[1] == "description" {
			fmt.Println("Description:", args[2])
		}
		os.Exit(0)
	}
	fmt.Println("args not found")
	os.Exit(1)
}
