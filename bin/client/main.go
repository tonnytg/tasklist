package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	ClientHelp = `pass help as argument to get help.
./taskclient help`
	AddHelp  = `Help Add!`
	ErrUsage = "args not found"
)

func main() {
	fmt.Println("Client Side")

	if err := processArgs(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func processArgs(args []string) error {
	if len(args) == 0 {
		fmt.Println(ClientHelp)
		return nil
	}

	switch args[0] {
	case "help":
		fmt.Println(AddHelp)
		return nil
	case "add":
		return handleAdd(args[1:])
	default:
		return errors.New(ErrUsage)
	}
}

func handleAdd(args []string) error {
	if len(args) == 0 {
		fmt.Println(AddHelp)
		return errors.New("missing arguments for add")
	}

	fmt.Println("Add")
	fmt.Println("lenArgs:", len(args))
	fmt.Println("Title:", args[0])

	if len(args) > 1 && (args[1] == "desc" || args[1] == "description") {
		if len(args) < 3 {
			return errors.New("missing description value")
		}
		fmt.Println("Description:", args[2])
	}
	return nil
}
