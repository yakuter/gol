package pwd

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	PmdHelp = "pwd"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:            PmdHelp,
		HelpName:        PmdHelp,
		Action:          Action,
		Usage:           `displays help messages.`,
		Description:     `Display help messages.`,
		SkipFlagParsing: true,
		HideHelp:        true,
		HideHelpCommand: true,
	}
}

func Action(c *cli.Context) error {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	return err
}
