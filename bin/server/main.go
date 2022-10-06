package main

import (
	"flag"
	"fmt"
	"github.com/tonnytg/tasklist/internal/daemon"
	"github.com/tonnytg/tasklist/pkg/api"
	"os"
)

const ServerHelp string = `
HELP Server!
`

func main() {
	fmt.Println("Server Side")

	if len(os.Args) < 2 {
		fmt.Println(ServerHelp)
		fmt.Println("Example: ./script --option api")
		os.Exit(0)
	}

	option := flag.String("option", "", "--option <VALUE>")
	help := flag.Bool("help", false, "--help")

	flag.Parse()

	if *help != false {
		fmt.Println(ServerHelp)
	}

	switch *option {
	case "api":
		// Start API Server
		api.Start()
	case "daemon":
		// Start Daemon Server
		daemon.Start()
	}

}
