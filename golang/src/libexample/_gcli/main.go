package main

import (
	"os"

	"github.com/codegangsta/cli"
)

//go:generate gcli new -command=test:"Test new task" example
func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "kyokomi"
	app.Email = ""
	app.Usage = ""

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
