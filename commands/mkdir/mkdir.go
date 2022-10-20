package mkdir

import (
	"os"

	"github.com/urfave/cli/v2"
)

const (
	CmdMkdir = "mkdir"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:            CmdMkdir,
		HelpName:        CmdMkdir,
		Action:          Action,
		Usage:           `createa a directory.`,
		Description:     `Creates a directory.`,
		SkipFlagParsing: true,
		HideHelp:        true,
		HideHelpCommand: true,
	}
}

func Action(c *cli.Context) error {
	var err error
	if err := os.MkdirAll(c.Args().First(), os.ModePerm); err != nil {
		err = cli.ShowAppHelp(c)
	}
	return err
}
