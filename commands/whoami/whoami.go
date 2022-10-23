package whoami

import (
	"fmt"
	"log"
	"os/user"

	"github.com/urfave/cli/v2"
)

const (
	CmdWhoami = "whoami"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:        CmdWhoami,
		HelpName:    CmdWhoami,
		Action:      Action,
		ArgsUsage:   ` `,
		Usage:       `display effective user id.`,
		Description: `The whoami utility displays your effective user ID as a name.`,
	}
}

func Action(c *cli.Context) error {
	u, err := user.Current()
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(u.Username)

	return nil
}
