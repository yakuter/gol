package pwd_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"

	"github.com/yakuter/gol/commands/pwd"
)

func TestPwd(t *testing.T) {
	execName, err := os.Executable()
	require.NoError(t, err)

	app := &cli.App{
		Commands: []*cli.Command{
			pwd.Command(),
		},
	}

	testArgs := []string{execName, "pwd"}
	require.NoError(t, app.Run(testArgs))
}
