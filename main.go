package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "wgx"
	app.Version = Version
	app.Usage = ""
	app.Author = "knmkr"
	app.Email = "knmkr3gma@gmail.com"
	app.Commands = Commands
	app.Run(os.Args)
}

// Version will be set at compile time by `-ldflags "-X main.Version=${VERSION}"`
var Version = "HEAD"

var Commands = []cli.Command{
	commandRunServer,
}

var commandRunServer = cli.Command{
	Name:   "runserver",
	Usage:  "",
	Action: doRunServer,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "addr",
			Usage: "e.g. localhost:1323",
		},
	},
}
