package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/tonnytg/tasklist/internal/binary"
	"github.com/tonnytg/tasklist/internal/daemon"
	"github.com/tonnytg/tasklist/pkg/api"
)

const ManHelp string = `

HELP!
Hello friend maybe you need some help!

--help	show this menu

--option <VALUE>
	at this case VALUE can be:
		api
		daemon
	Example: ./script --option webserver

--add "TASK Title"`

func Cmd() {

	if len(os.Args) == 1 {
		fmt.Println(ManHelp)
		os.Exit(0)
	}

	option := flag.String("option", "", "--option <VALUE>")
	add := flag.String("add", "", "--add <VALUE>")
	help := flag.Bool("help", false, "--help")

	flag.Parse()

	if *help != false {
		fmt.Println(ManHelp)
	}

	if *add != "" {
		binary.Create()
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
