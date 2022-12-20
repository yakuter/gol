package wc

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"strings"
)

const (
	CmdWc = "wc"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:        CmdWc,
		HelpName:    CmdWc,
		Action:      Action,
		ArgsUsage:   ` `,
		Usage:       `find out the number of lines, word count, byte and characters count in the files specified in the file arguments`,
		Description: `The wc utility can be used to find out the number of lines, word count, byte and characters count in the files specified in the file arguments`,
		Flags:       nil,
	}
}

func Action(c *cli.Context) error {
	args := c.Args()
	fileName := args.First()
	if args.Len() < 1 {
		return cli.ShowAppHelp(c)
	}

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(file)
	lines, words, characters := 0, 0, 0

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		lines++
		words += len(strings.Fields(line))
		characters += len(line)
	}

	fmt.Printf("%8d%8d%8d %s\n", lines, words, characters, fileName)

	return nil
}
