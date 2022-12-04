package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"

	"github.com/urans/bing/pkg/banner"
	"github.com/urans/bing/pkg/dict"
)

var flags = []cli.Flag{
	&cli.IntFlag{
		Name:    "sentnum",
		Aliases: []string{"n"},
		Value:   5,
		Usage:   "number of example sentences",
	},
}

var action = func(ctx *cli.Context) error {
	if ctx.Args().Len() > 0 && ctx.Args().First() != "help" {
		err := dict.Query(ctx.Args().First(), ctx.Int("sentnum"))
		if err != nil {
			color.Red(err.Error())
		}
		return err
	}

	banner.Print()
	return cli.ShowAppHelp(ctx)
}

func main() {
	app := cli.NewApp()
	app.Usage = "A Command Line Dictionary for Geekers 📝"
	app.Flags = flags
	app.Action = action
	app.Commands = []*cli.Command{}
	app.UseShortOptionHandling = true
	app.Run(os.Args)
}
