package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)
}

func main() {
	app := cli.NewApp()
	app.Name = "awtk"
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
