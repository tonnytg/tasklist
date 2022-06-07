package cmd

import (
	"flag"
	"fmt"
)

const ManHelp string = `

HELP!
Hello friend maybe you need some help!

--help	show this menu

--option <VALUE>
	at this case VALUE can be:
		knife
		sword
		magic
	Example: ./script --option sword`



func Cmd() {

	option := flag.String("option", "", "--option <VALUE>")
	help := flag.Bool("help", false, "--help")

	flag.Parse()

	if *help != false {
		fmt.Println(ManHelp)
	}

	fmt.Println("option selected:", *option)
}
