package touch_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
	"github.com/yakuter/gol/commands/touch"
)

var waitTime = time.Millisecond * 200
var expectedTimeDelta = time.Millisecond * 10

func TestTouch(t *testing.T) {
	t.Parallel()

	execName, err := os.Executable()
	require.NoError(t, err)

	app := &cli.App{
		Commands: []*cli.Command{
			touch.Command(),
		},
	}

	tests := []struct {
		name     string
		args     []string
		testFunc func()
	}{
		{
			name: "only touch arg",
			args: []string{execName, "touch"},
		},
		{
			name: "create single file",
			args: []string{execName, "touch", "hello"},
			testFunc: func() {
				require.FileExists(t, "hello")
				os.Remove("hello") // Clean
			},
		},
		{
			name: "create multiple files",
			args: []string{execName, "touch", "one", "two"},
			testFunc: func() {
				require.FileExists(t, "one")
				require.FileExists(t, "two")
				// Clean
				os.Remove("one")
				os.Remove("two")
			},
		},
		{
			name: "change time of existing file",
			args: []string{execName, "touch", "hello"},
			testFunc: func() {
				require.FileExists(t, "hello")

				first, err := os.Stat("hello")
				require.NoError(t, err)

				time.Sleep(waitTime)

				require.NoError(t, app.Run([]string{execName, "touch", "hello"}))
				secondTime := time.Now()
				second, err := os.Stat("hello")
				require.NoError(t, err)

				require.WithinDuration(t, secondTime, second.ModTime(), expectedTimeDelta)
				require.WithinDuration(t, first.ModTime(), second.ModTime(), expectedTimeDelta+waitTime)

				// Clean
				os.Remove("hello")
			},
		},
		// {
		// 	name: "use --no-create flag",
		// 	args: []string{execName, "touch", "--no-create", "hello"},
		// 	testFunc: func() {
		// 		require.NoFileExists(t, "hello")
		// 	},
		// },
		// {
		// 	name: "use -c alias",
		// 	args: []string{execName, "touch", "-c", "hello"},
		// 	testFunc: func() {
		// 		require.NoFileExists(t, "hello")
		// 	},
		// },
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.NoError(t, app.Run(test.args))

			if test.testFunc != nil {
				test.testFunc()
			}
		})
	}
}
