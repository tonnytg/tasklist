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

--add "TASK Title"
--description "TASK Description"
--list`

func Cmd() {

	if len(os.Args) == 1 {
		fmt.Println(ManHelp)
		os.Exit(0)
	}

	option := flag.String("option", "", "--option <VALUE>")
	add := flag.String("add", "", "--add <VALUE>")
	description := flag.String("description", "", "--description <VALUE>")
	list := flag.Bool("list", false, "--list")
	help := flag.Bool("help", false, "--help")

	flag.Parse()

	if *help != false {
		fmt.Println(ManHelp)
	}

	if *add != "" {
		binary.Create(*add, *description)
	}

	if *list != false {
		binary.List()
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
