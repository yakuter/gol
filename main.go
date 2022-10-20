package main

import (
	"io"
	"log"
	"os"

	"github.com/yakuter/gol/commands/echo"
	"github.com/yakuter/gol/commands/help"

	"github.com/urfave/cli/v2"
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
	}
}
