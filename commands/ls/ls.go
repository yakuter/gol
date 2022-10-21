package ls

import (
	"fmt"
	"io/ioutil"
	"log"

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
	args := c.Args()
	var files,err = ioutil.ReadDir(args.First())

	if args.Len() == 0 {
		files,err = ioutil.ReadDir("./")
	}

	if err != nil {
		log.Println(err)
		return err
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

	return nil
}
