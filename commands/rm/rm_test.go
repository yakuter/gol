package rm_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"

	"github.com/yakuter/gol/commands/rm"
)

func TestRm_WithFiles(t *testing.T) {
	execName, err := os.Executable()
	require.NoError(t, err)

	app := &cli.App{
		Commands: []*cli.Command{
			rm.Command(),
		},
	}

	file1, err := os.CreateTemp("./", "test")
	require.NoError(t, err)
	file2, err := os.CreateTemp("./", "test")
	require.NoError(t, err)

	time.Sleep(1 * time.Second)
	testArgs := []string{execName, "rm", file1.Name(), file2.Name()}
	require.NoError(t, app.Run(testArgs))
	require.False(t, fileExists(file1.Name()))
	require.False(t, fileExists(file2.Name()))
}

func TestRm_WithDirectories(t *testing.T) {
	execName, err := os.Executable()
	require.NoError(t, err)

	app := &cli.App{
		Commands: []*cli.Command{
			rm.Command(),
		},
	}

	dir1, err := os.MkdirTemp("./", "test")
	require.NoError(t, err)
	dir2, err := os.MkdirTemp("./"+dir1, "test")
	require.NoError(t, err)
	file1, err := os.CreateTemp("./"+dir1, "test")
	require.NoError(t, err)
	file2, err := os.CreateTemp("./"+dir2, "test")
	require.NoError(t, err)

	testArgs := []string{execName, "rm", "-r", dir1}
	require.NoError(t, app.Run(testArgs))

	require.False(t, fileExists(file1.Name()))
	require.False(t, fileExists(file2.Name()))
	require.False(t, fileExists(dir1))
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
