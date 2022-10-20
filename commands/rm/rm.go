package rm

import (
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

const (
	CmdRm = "rm"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:        CmdRm,
		HelpName:    CmdRm,
		Action:      Action,
		Usage:       `removes file or directory.`,
		Description: `removes file or directory.`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "force",
				Aliases:     []string{"f"},
				Usage:       "force remove",
				Required:    false,
				DefaultText: "false",
			},
			&cli.BoolFlag{
				Name:        "recursive",
				Aliases:     []string{"r"},
				Usage:       "remove directories and their contents recursively",
				Required:    false,
				DefaultText: "false",
			},
		},
		HideHelp:        true,
		HideHelpCommand: true,
	}
}

func Action(c *cli.Context) error {
	if c.Bool("recursive") {
		if err := removeDirectories(c.Args().Slice()); err != nil {
			cli.ShowAppHelp(c)
			return err
		}
	} else {
		if err := removeFiles(c.Args().Slice()); err != nil {
			cli.ShowAppHelp(c)
			return err
		}
	}

	return nil
}

func removeFiles(files []string) error {
	for _, file := range files {
		if err := removeFile(file, false); err != nil {
			return err
		}
	}

	return nil
}

func removeDirectories(dirs []string) error {
	for _, dir := range dirs {
		if err := removeDirectoryWithContents(dir); err != nil {
			return err
		}
	}

	return nil
}

func removeDirectoryWithContents(dir string) error {
	// Open the file and get the file info
	file, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	// If the file is not a directory, remove it as a file
	if !stat.IsDir() {
		return removeFile(dir, false)
	}

	// Get all files in the directory
	names, err := file.Readdirnames(100)
	if err != nil {
		return err
	}

	for _, name := range names {
		stat, err := os.Stat(filepath.Join(dir, name))
		if err != nil {
			return err
		}

		// If the file is a directory, remove it recursively
		if stat.IsDir() {
			if err := removeDirectoryWithContents(filepath.Join(dir, name)); err != nil {
				return err
			}
		}

		if err := removeFile(filepath.Join(dir, name), true); err != nil {
			return err
		}
	}

	// Remove top directory
	if err := removeFile(dir, true); err != nil {
		return err
	}

	return nil
}

func removeFile(file string, isDir bool) error {
	if isDir {
		return os.RemoveAll(file)
	}

	if err := os.Remove(file); err != nil {
		return err
	}

	return nil
}
