package grep

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

const (
	CmdGrep = "grep"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:        CmdGrep,
		HelpName:    CmdGrep,
		Action:      Action,
		ArgsUsage:   "<pattern> [file]",
		Usage:       `print lines that match patterns`,
		Description: `The grep utility searches the specified pattern in the specified file or STDIN`,
	}
}

var (
	sprintfRed = color.New(color.FgRed).SprintFunc()
)

func Action(c *cli.Context) error {
	args := c.Args()
	if args.Len() == 0 || args.Len() > 2 {
		return fmt.Errorf("invalid number of args, usage: grep %s", c.Command.ArgsUsage)
	}

	var in io.Reader = os.Stdin
	if args.Len() == 2 {
		path := args.Get(1)
		f, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open the input file '%s', error: %v", path, err)
		}
		defer f.Close()
		in = f
	}

	re := regexp.MustCompilePOSIX(args.First())
	sc := bufio.NewScanner(in)
	for sc.Scan() {
		line := sc.Text()
		index := re.FindIndex([]byte(line))
		if index != nil {
			fmt.Printf("%s%s%s\n", line[0:index[0]], sprintfRed(line[index[0]:index[1]]), line[index[1]:])
		}
	}
	if err := sc.Err(); err != nil {
		return fmt.Errorf("failed to read the input, error: %v", err)
	}

	return nil
}
