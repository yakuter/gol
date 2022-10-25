package mv

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	CmdMv = "mv"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:            CmdMv,
		HelpName:        CmdMv,
		Action:          Action,
		ArgsUsage:       ` `,
		Usage:           `move a file or directories from one place to another`,
		Description:     `The mv utility  is used to move one or more files or directories from one place to another in a file system`,
		SkipFlagParsing: true,
		HideHelp:        true,
		HideHelpCommand: true,
	}
}

func Action(c *cli.Context) error {
	var err error
	if err = os.Rename(c.Args().First(), c.Args().Get(2)); err != nil {
		err = cli.ShowAppHelp(c)
		log.Println(err)
	}
	return err
}
