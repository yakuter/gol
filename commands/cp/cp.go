package cp

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

const (
	CmdCp = "cp"
)

var (
	pwd string
)

func Command() *cli.Command {
	return &cli.Command{
		Name:     CmdCp,
		HelpName: CmdCp,
		Action:   Action,
		ArgsUsage: `
	 gol cp -r source_dir target_dir
	 gol cp source_file target_file`,
		Usage:       `Copy SOURCE to TARGET`,
		Description: `Copy SOURCE to TARGET`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "r",
				Value: false,
				Usage: "gol cp -r source_dir target_dir",
			},
		},
	}
}

func Action(c *cli.Context) error {
	args := c.Args().Slice()

	// minimum 2 paths are required
	if len(args) < 2 {
		return errors.New("missing file operands")
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("cannot %v", err)
		os.Exit(1)
	}

	pwd = wd

	// the last argument is the target path
	dest := abs(args[len(args)-1])
	destStat, err := os.Stat(dest)
	if err == nil {
		// target is exists
		if destStat.Mode().IsRegular() {
			// and its not a dir
			return errors.New("target file is is already exists")
		}
	}

	for _, path := range args[:len(args)-1] {
		path = abs(path)

		// get path stats
		stat, err := os.Stat(path)
		if err != nil {
			return fmt.Errorf("cannot %w", err)
		}

		target := dest
		if destStat != nil && destStat.IsDir() {
			target = fmt.Sprintf("%s/%s", dest, filepath.Base(path))
		}

		if stat.Mode().IsRegular() {
			if err := cpFile(c, path, target, stat.Size()); err != nil {
				return err
			}
		} else if stat.IsDir() {
			if err := cpDir(c, path, target); err != nil {
				return err
			}
		}
	}

	return nil
}

func cpFile(c *cli.Context, path, dest string, bufSize int64) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not open %s, %w", path, err)
	}
	defer f.Close()

	target, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("could not create %s, %w", dest, err)
	}
	defer target.Close()

	buf := make([]byte, bufSize)
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			return fmt.Errorf("could not read file %s, %w", path, err)
		}
		if n == 0 {
			break
		}

		if _, err := target.Write(buf[:n]); err != nil {
			return err
		}
	}

	return nil
}

func cpDir(c *cli.Context, path, dest string) error {
	pathStat, err := os.Stat(path)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dest, pathStat.Mode().Perm()); err != nil {
		return err
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			if err := cpDir(c, path+"/"+f.Name(), dest+"/"+f.Name()); err != nil {
				return err
			}
		} else {
			fStat, err := os.Stat(path + "/" + f.Name())
			if err != nil {
				return err
			}
			if err := cpFile(c, path+"/"+f.Name(), dest+"/"+f.Name(), fStat.Size()); err != nil {
				return err
			}
		}
	}

	return nil
}

func abs(path string) string {
	// if path is not an absolute path
	// prefix it with current working dir
	if !filepath.IsAbs(path) {
		path = filepath.Join(pwd, path)
	}

	return path
}
