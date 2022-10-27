package mv_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
	"github.com/yakuter/gol/commands/mv"
)

func TestHelp(t *testing.T) {
	execName, err := os.Executable()
	require.NoError(t, err)

	app := &cli.App{
		Commands: []*cli.Command{
			mv.Command(),
		},
	}

	testArgs := []string{execName, "mv"}
	require.NoError(t, app.Run(testArgs))
}
