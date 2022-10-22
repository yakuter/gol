package cp_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
	"github.com/yakuter/gol/commands/cp"
)

func TestCp(t *testing.T) {
	execName, err := os.Executable()
	require.NoError(t, err)

	app := &cli.App{
		Commands: []*cli.Command{
			cp.Command(),
		},
	}

	testArgs := []string{execName, "cp", "source", "target"}
	require.NoError(t, app.Run(testArgs))
}
