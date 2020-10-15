package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/magicalbanana/bentobox/dirls"
	"github.com/magicalbanana/bentobox/strcomp"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:  "dirls",
			Usage: "Recursively walks a directory and lists the contents with the file size",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "sort-by",
					Value: "asc",
					Usage: "sort by either asc or desc",
				},
				cli.StringFlag{
					Name:  "output",
					Value: "simple",
					Usage: "Print either as a flat directory or tree directory, vlaues are: simple | tree",
				},
			},
			Action: func(c *cli.Context) error {
				cli.CommandHelpTemplate = strings.Replace(cli.CommandHelpTemplate, "[arguments...]", "<dir>", -1)
				dir := c.Args().First()
				if dir == "" {
					cli.ShowCommandHelp(c, "dirls")
					os.Exit(1)
				}
				ff, node, err := dirls.DirLs(dir)
				if err != nil {
					return err
				}
				sortByStr := c.String("sort-by")
				sortBy := dirls.ASC
				if sortByStr == "desc" {
					sortBy = dirls.DESC
				}

				output := c.String("output")
				if output == "tree" {
					dirls.PrintTree(node, "")
					return nil
				}

				dirls.PrintFiles(ff, sortBy)
				return nil
			},
		},

		{
			Name:  "strcomp",
			Usage: "Compress a string",
			Action: func(c *cli.Context) error {
				cli.CommandHelpTemplate = strings.Replace(cli.CommandHelpTemplate, "[arguments...]", "<word to compress>", -1)
				str := c.Args().First()
				if str == "" {
					cli.ShowCommandHelp(c, "strcomp")
					os.Exit(1)
				}
				fmt.Println(strcomp.Compress(str))

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
