package rm

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

const (
	CmdRm                = "rm"
	Unwritable_File_Perm = 0444
)

func Command() *cli.Command {
	return &cli.Command{
		Name:        CmdRm,
		HelpName:    CmdRm,
		Action:      Action,
		Usage:       `removes files or directories.`,
		Description: `removes files or directories.`,
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
				Aliases:     []string{"R"},
				Usage:       "remove directories and their contents recursively",
				Required:    false,
				DefaultText: "false",
			},
			// TODO: implement forcefully recursive
			// &cli.BoolFlag{
			// 	Name:        "forcefully_recursive",
			// 	Aliases:     []string{"rf"},
			// 	Usage:       "remove directories and their contents forcefully recursively",
			// 	Required:    false,
			// 	DefaultText: "false",
			// },
		},
		HideHelp:        true,
		HideHelpCommand: true,
	}
}

func Action(c *cli.Context) error {
	force := c.Bool("force") || c.Bool("forcefully_recursive")
	recursive := c.Bool("recursive") || c.Bool("forcefully_recursive")

	if recursive {
		if err := removeDirectories(c.Args().Slice(), force); err != nil {
			cli.ShowAppHelp(c)
			return err
		}
	} else {
		if err := removeFiles(c.Args().Slice(), force); err != nil {
			cli.ShowAppHelp(c)
			return err
		}
	}

	return nil
}

func removeFiles(files []string, force bool) error {
	for _, file := range files {
		if err := removeFile(file, false, force); err != nil {
			return err
		}
	}

	return nil
}

func removeDirectories(dirs []string, force bool) error {
	for _, dir := range dirs {
		if err := removeDirectoryWithContents(dir, force); err != nil {
			return err
		}
	}

	return nil
}

func removeDirectoryWithContents(dir string, force bool) error {
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
		return removeFile(dir, false, force)
	}

	// Get all files in the directory
	fileNames, err := file.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, name := range fileNames {
		stat, err := os.Stat(filepath.Join(dir, name))
		if err != nil {
			return err
		}

		// If the file is a directory, remove it recursively
		if stat.IsDir() {
			if err := removeDirectoryWithContents(filepath.Join(dir, name), force); err != nil {
				return err
			}
		}
	}

	// Remove top directory
	if err := removeFile(dir, true, force); err != nil {
		return err
	}

	return nil
}

func removeFile(file string, isDir bool, force bool) error {
	if !force {
		stat, err := os.Stat(file)
		if err != nil {
			return err
		}

		// if unwritable, ask to user if he wants to force remove
		if stat.Mode().Perm() == Unwritable_File_Perm {
			var answer string
			print("rm: remove write-protected regular file '" + file + "'?(y/n) ")
			if fmt.Scan(&answer); answer != "y" {
				return nil
			}
		}
	}

	if isDir {
		return os.RemoveAll(file)
	}

	if err := os.Remove(file); err != nil {
		return err
	}

	return nil
}
