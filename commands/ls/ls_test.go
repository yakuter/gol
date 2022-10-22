package ls_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"

	"github.com/yakuter/gol/commands/ls"
)

func TestLs(t *testing.T) {
	execName, err := os.Executable()
	require.NoError(t, err)

	app := &cli.App{
		Commands: []*cli.Command{
			ls.Command(),
		},
	}

	testArgs := []string{execName, "ls"}
	require.NoError(t, app.Run(testArgs))
}
