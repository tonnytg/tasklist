package cmd

import (
	"flag"
	"fmt"
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
	Example: ./script --option webserver`



func Cmd() {

	option := flag.String("option", "", "--option <VALUE>")
	help := flag.Bool("help", false, "--help")

	flag.Parse()

	if *help != false {
		fmt.Println(ManHelp)
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
