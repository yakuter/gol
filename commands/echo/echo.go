package echo

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

const (
	CmdEcho = "echo"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:        CmdEcho,
		HelpName:    CmdEcho,
		Action:      Action,
		ArgsUsage:   ` `,
		Usage:       `write arguments to the standard output.`,
		Description: `The echo utility writes any specified text to the standard output.`,
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
	args := c.Args()
	if args.Len() == 0 {
		return nil
	}

	output := strings.Join(args.Slice(), " ")

	fmt.Println(output)

	return nil
}
