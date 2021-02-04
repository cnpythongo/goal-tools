package main

import (
	"fmt"
	"github.com/cnpythongo/goal-tools/start"
	"github.com/urfave/cli"
	"os"
	"runtime"
)

var (
	Version = "0.0.1"

	commands = []cli.Command{
		// goal-tools startproject {projectname}
		{
			Name:   "startproject",
			Usage:  "create start project from github.com/cnpythongo/goal-helper template",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "dir,d",
					Usage: "create new project in target directory",
				},
				cli.StringFlag{
					Name:  "app,a",
					Usage: "app service name for project",
				},
			},
			Action: start.NewProject,
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Usage = "a cli tool to generate web framework code"
	app.Version = fmt.Sprintf("%s %s/%s", Version, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		fmt.Println("error:", err)
	}
}
