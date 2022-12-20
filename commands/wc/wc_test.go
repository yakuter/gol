package wc_test

import (
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
	"github.com/yakuter/gol/commands/wc"
	"os"
	"testing"
)

func TestEcho(t *testing.T) {
	execName, err := os.Executable()
	require.NoError(t, err)

	app := &cli.App{
		Commands: []*cli.Command{
			wc.Command(),
		},
	}

	testArgs := []string{execName, "wc"}
	require.NoError(t, app.Run(testArgs))
}
