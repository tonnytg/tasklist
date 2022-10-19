package main

import (
	"flag"
	"fmt"
	"os"
)

const ClientHelp string = `
Help Client!
`

func main() {
	fmt.Println("Client Side")

	if len(os.Args) < 2 {
		fmt.Println(ClientHelp)
		fmt.Println("Example: ./script --option api")
		os.Exit(0)
	}

	//add := flag.String("add", "", "--add <VALUE>")
	//description := flag.String("description", "", "--description <VALUE>")
	//list := flag.Bool("list", false, "--list")
	help := flag.Bool("help", false, "--help")

	flag.Parse()

	if *help != false {
		fmt.Println(ClientHelp)
	}
	//
	//if *add != "" {
	//	binary.Create(*add, *description)
	//}
	//
	//if *list != false {
	//	binary.List()
	//}

}
