package main

import (
	"fmt"
	"os"

	"github.com/docker/docker/pkg/homedir"
	"github.com/urfave/cli"
)

var Version string

func main() {
	if Version == "" {
		Version = "unset/develop"
	}

	app := cli.NewApp()
	app.Name = "Context"
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "dir",
			Value:  fmt.Sprintf("%s/.context", homedir.Get()),
			Usage:  "Directory to manage contexts",
			EnvVar: "CTX_DIR",
		},
	}

	app.Before = Before

	app.Commands = []cli.Command{
		{
			Name:      "edit",
			Usage:     "Edit a context",
			ArgsUsage: "CONTEXT",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "editor, e",
					Value:  "vim",
					Usage:  "Editor used to edit the context",
					EnvVar: "CTX_EDITOR",
				},
			},
			Action: Edit,
		},
		{
			Name:      "delete",
			Aliases:   []string{"del"},
			ArgsUsage: "CONTEXT...",
			Usage:     "Delete context(s)",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "force, f",
					Usage: "Disable prompting before removal",
				},
			},
			Action: Delete,
		},
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "List available contexts",
			Action:  List,
		},
		{
			Name:      "switch",
			Usage:     "Switch to a context",
			ArgsUsage: "CONTEXT",
			Action:    Switch,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
