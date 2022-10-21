package ls

import (
	"fmt"
	"io/ioutil"

	"github.com/urfave/cli/v2"
)

const (
	CmdLs = "ls"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:        CmdLs,
		HelpName:    CmdLs,
		Action:      Action,
		ArgsUsage:   ` `,
		Usage:       `write arguments to the standard output.`,
		Description: `The ls utility lists directory contents of files and directories`,
		Flags:       Flags(),
	}
}

func Flags() []cli.Flag {
	return []cli.Flag{
		// Example Flags
		/*
			&cli.StringFlag{
				Name:     flagOut,
				Usage:    "Output file path",
				Required: false,
			},
			&cli.UintFlag{
				Name:     flagBits,
				Usage:    "Number of bits",
				Required: true,
			},
			&cli.BoolFlag{
				Name:        flagWithPublic,
				Usage:       "Export public key with private key",
				Required:    false,
				DefaultText: "false",
			},
		*/
	}
}

func Action(c *cli.Context) error {
	// args := c.Args()
	files, err := ioutil.ReadDir("./")

	if err != nil {
		fmt.Println(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

	return nil
}
