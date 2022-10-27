package main

import (
	"io"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/yakuter/gol/commands/cp"
	"github.com/yakuter/gol/commands/echo"
	"github.com/yakuter/gol/commands/grep"
	"github.com/yakuter/gol/commands/help"
	"github.com/yakuter/gol/commands/ls"
	"github.com/yakuter/gol/commands/mkdir"
	"github.com/yakuter/gol/commands/pwd"
	"github.com/yakuter/gol/commands/whoami"
)

var Version = "v1.0.0"

func main() {
	app := &cli.App{
		Name:     "Gol",
		Usage:    "Go implementation of Linux commands",
		Commands: Commands(os.Stdin),
		Version:  Version,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func Commands(reader io.Reader) []*cli.Command {
	return []*cli.Command{
		help.Command(),
		echo.Command(),
		pwd.Command(),
		mkdir.Command(),
		whoami.Command(),
		grep.Command(),
		ls.Command(),
		cp.Command(),
	}
}
