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

	tmpDir := os.TempDir()

	p, err := os.MkdirTemp("", "source_dir")
	if err != nil {
		t.Fatalf("could not create temp dir, %v", err)
	}

	f, err := os.CreateTemp(p, "source")
	if err != nil {
		t.Fatalf("could not create temp file, %v", err)
	}
	defer os.Remove(f.Name()) // clean up

	testArgs := []string{execName, "cp", f.Name(), tmpDir + "/target"}
	require.NoError(t, app.Run(testArgs))

	testArgs = []string{execName, "cp", "-r", p, tmpDir + "/target_dir"}
	require.NoError(t, app.Run(testArgs))

	t.Cleanup(func() {
		os.Remove(tmpDir + "/target")
		os.RemoveAll(tmpDir + "/target_dir")
	})
}
