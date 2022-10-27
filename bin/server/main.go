package main

import (
	"flag"
	"fmt"
	"github.com/tonnytg/tasklist/pkg/api"
)

const ServerHelp string = `
HELP Server!
`

func main() {
	fmt.Println("Server Side")

	option := flag.String("option", "api", "--option <VALUE>")
	help := flag.Bool("help", false, "--help")

	flag.Parse()

	if *help != false {
		fmt.Println(ServerHelp)
	}

	switch *option {
	case "api":
		// Start API Server
		api.Start()
	}
}
