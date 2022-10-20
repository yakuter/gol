package touch

import (
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

const (
	CmdTouch = "touch"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:        CmdTouch,
		HelpName:    CmdTouch,
		Action:      Action,
		ArgsUsage:   ` `,
		Usage:       `creates empty files or change file timestamps`,
		Description: `Updates the access and modification times of each file, creates an empty file if not exists`,
		Flags:       Flags(),
	}
}

func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:     "no-create",
			Aliases:  []string{"c"},
			Usage:    "do not create new file",
			Required: false,
		},
	}
}

func Action(c *cli.Context) error {
	args := c.Args()
	if c.Args().Len() == 0 {
		return cli.ShowAppHelp(c)
	}

	if !c.Bool("no-create") {
		for _, arg := range args.Slice() {
			_, err := os.Create(arg)
			if err != nil {
				return cli.Exit("error occurred while creating file", 5)
			}
		}
	}

	now := time.Now()
	for _, arg := range args.Slice() {
		err := os.Chtimes(arg, now, now)

		if os.IsNotExist(err) {
			notExistMsg := "no such file exist: " + arg
			return cli.Exit(notExistMsg, 2)
		}

		if err != nil {
			return cli.Exit("error occurred while changing times", 5)
		}
	}

	return nil
}
