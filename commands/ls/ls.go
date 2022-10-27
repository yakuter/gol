package ls

import (
	"fmt"
	"os"

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
		Usage:       `The ls utility lists directory contents of files and directories`,
		Description: `The ls utility lists directory contents of files and directories`,
		Flags:       Flags(),
	}
}

func Flags() []cli.Flag {
	return []cli.Flag{
		// TODO: Add flags
	}
}

func Action(c *cli.Context) error {
	path := c.Args().First()
	if path == "" {
		path = "./"
	}
 	files, err := os.ReadDir(path)
 	if err != nil {
 		return fmt.Errorf("read dir %q error: %w", path, err)
 	}
 	for _, f := range files {
 		fmt.Println(f.Name())
 	}
 	return nil
 }
