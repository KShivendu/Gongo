package main

import (
	"fmt"
	"gongo/utils"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

// TODO: Add CLI that generates main.go

func main() {
	app := (&cli.App{
		Name:  "gongo",
		Usage: "A CLI for Gongo",
		Commands: []*cli.Command{
			{
				Name:    "startproject",
				Aliases: []string{"sp"},
				Usage:   "Start a Gongo project",
				Action: func(c *cli.Context) error {
					var projectName = c.Args().First()
					var target = "."
					// TODO: Check that target directory exists
					utils.CopyDirectory("boilerplate/project", target+"/"+projectName)
					fmt.Println("Created project!")
					return nil
				},
			},
			{
				Name:    "startapp",
				Aliases: []string{"sa"},
				Usage:   "Start a Gongo app in the project",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "runserver",
				Aliases: []string{"r"},
				Usage:   "Run the Gongo server",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			return nil
		},
	})

	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
