package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kyokomi-sandbox/go-sandbox/libexample/_gcli/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{

	{
		Name:   "test",
		Usage:  "Test new task",
		Action: command.CmdTest,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
