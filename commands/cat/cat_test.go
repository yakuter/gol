package cat

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestStrSliceToMap(t *testing.T) {
	temp := []string{"a", "b", "c"}
	expectedReturn := map[string]bool{"a": true, "b": true, "c": true}
	result := strSliceToMap(temp)

	assert.Equal(t, true, result[temp[0]])
	assert.Equal(t, false, result["false"])

	assert.Equal(t, expectedReturn, result)
}

func TestInsert(t *testing.T) {
	temp := []byte{'a', 'b', 'c'}
	expectedReturn := []byte{'d', 'a', 'b', 'c'}
	result := insert(temp, 0, 'd')

	assert.Equal(t, expectedReturn, result)
}

func TestRemove(t *testing.T) {
	temp := []byte{'a', 'b', 'c'}
	expectedReturn := []byte{'b', 'c'}
	result := remove(temp, 0)

	assert.Equal(t, expectedReturn, result)
}

func TestDisplayNoNPrt(t *testing.T) {
	temp := []byte{1, 2, 3, 48, 127, 150, 160, 255, '\t', '\n', '\r'}
	expectedReturn := []byte{94, 65, 94, 66, 94, 67, 48, 94, 63, 77, 45, 94, 86, 77, 45, 32, 77, 45, 94, 63, 9, 10, 94, 77}
	result := displayNoNPrt(temp, false)
	assert.Equal(t, expectedReturn, result)
}

func TestDisplayEndOfLine(t *testing.T) {
	testCases := [][]byte{
		{'\n'},
		{'a', 'b', 'c', '\r', '\n'},
		{'a', 'b', 'c', '\n'},
	}

	testCaseOneResult := displayEndOfLine(testCases[0])
	testCaseTwoResult := displayEndOfLine(testCases[1])
	testCaseThreeResult := displayEndOfLine(testCases[2])

	assert.Equal(t, []byte{'$', '\n'}, testCaseOneResult)
	assert.Equal(t, []byte{'a', 'b', 'c', '$', 'M', '^', '\r', '\n'}, testCaseTwoResult)
	assert.Equal(t, []byte{'a', 'b', 'c', '$', '\n'}, testCaseThreeResult)
}

func TestDisplayTabCharacter(t *testing.T) {
	testCases := [][]byte{
		{'a', '\t', 'b', 'c', '\n'},
		{'a', 'b', 'c', '\n'},
	}

	testCaseOneResult := displayTabCharacter(testCases[0])
	testCaseTwoResult := displayTabCharacter(testCases[1])

	assert.Equal(t, []byte{'a', '^', 'I', 'b', 'c', '\n'}, testCaseOneResult)
	assert.Equal(t, []byte{'a', 'b', 'c', '\n'}, testCaseTwoResult)
}

func TestAddNumberForNonEmptyLine(t *testing.T) {
	testCases := [][]byte{
		{' '},
		{'a'},
	}

	testCasesOneCounter := 0
	testCasesTwoCounter := 1

	testCaseOneResult := addNumberForNonEmptyLine(testCases[0], &testCasesOneCounter)
	testCaseTwoResult := addNumberForNonEmptyLine(testCases[1], &testCasesTwoCounter)

	assert.Equal(t, []byte{' '}, testCaseOneResult)
	assert.Equal(t, []byte{49, 32, 'a'}, testCaseTwoResult)
}

func TestAddNumberForLine(t *testing.T) {
	testCases := [][]byte{
		{' '},
		{'a'},
	}

	testCasesOneCounter := 0
	testCasesTwoCounter := 1

	testCaseOneResult := addNumberForLine(testCases[0], &testCasesOneCounter)
	testCaseTwoResult := addNumberForLine(testCases[1], &testCasesTwoCounter)

	assert.Equal(t, []byte{48, 32, ' '}, testCaseOneResult)
	assert.Equal(t, []byte{49, 32, 'a'}, testCaseTwoResult)
}

type TaseCase struct {
	name           string
	flags          []string
	expectedOutput []byte
	expectedErr    error
}

func TestCatBasic(t *testing.T) {
	execName, err := os.Executable()
	assert.NoError(t, err)

	app := &cli.App{
		Commands: []*cli.Command{
			Command(),
		},
	}

	testFileName := "test_file_" + fmt.Sprint(time.Now().Unix())
	testFileContent := "test line 1 \n \ttest tab\n \ntest line 2\n test line 3 \r\n test line 4 işöğ"

	err = createTestFile(testFileName, testFileContent)
	if err != nil {
		t.Error("cannot create test file")
	}

	defer removeTestFile(testFileName)

	testCases := []TaseCase{
		{name: "basic", flags: []string{execName, "cat", testFileName}, expectedOutput: []byte(testFileContent)},
		{name: "basic not found", flags: []string{execName, "cat", ""}, expectedOutput: []byte{}, expectedErr: errors.New("no such file")},
		{
			name:           "-b flag",
			flags:          []string{execName, "cat", testFileName, "-b"},
			expectedOutput: []byte{49, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 49, 32, 10, 50, 32, 32, 9, 116, 101, 115, 116, 32, 116, 97, 98, 10, 32, 10, 51, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 50, 10, 52, 32, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 51, 32, 13, 10, 53, 32, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 52, 32, 105, 197, 159, 195, 182, 196, 159},
		},
		{
			name:           "-A flag",
			flags:          []string{execName, "cat", testFileName, "-A"},
			expectedOutput: []byte{116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 49, 32, 36, 10, 32, 94, 73, 116, 101, 115, 116, 32, 116, 97, 98, 36, 10, 32, 36, 10, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 50, 36, 10, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 51, 32, 36, 77, 94, 94, 77, 10, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 52, 32, 105, 77, 45, 69, 77, 45, 94, 95, 77, 45, 67, 77, 45, 54, 77, 45, 68, 77, 45, 94, 95},
		},
		{
			name:           "-e flag",
			flags:          []string{execName, "cat", testFileName, "-e"},
			expectedOutput: []byte{116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 49, 32, 36, 10, 32, 9, 116, 101, 115, 116, 32, 116, 97, 98, 36, 10, 32, 36, 10, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 50, 36, 10, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 51, 32, 36, 77, 94, 94, 77, 10, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 52, 32, 105, 77, 45, 69, 77, 45, 94, 95, 77, 45, 67, 77, 45, 54, 77, 45, 68, 77, 45, 94, 95},
		},
		{
			name:           "-E flag",
			flags:          []string{execName, "cat", testFileName, "-E"},
			expectedOutput: []byte{116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 49, 32, 36, 10, 32, 9, 116, 101, 115, 116, 32, 116, 97, 98, 36, 10, 32, 36, 10, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 50, 36, 10, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 51, 32, 36, 77, 94, 13, 10, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 52, 32, 105, 197, 159, 195, 182, 196, 159},
		},
		{
			name:           "-n flag",
			flags:          []string{execName, "cat", testFileName, "-b"},
			expectedOutput: []byte{49, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 49, 32, 10, 50, 32, 32, 9, 116, 101, 115, 116, 32, 116, 97, 98, 10, 32, 10, 51, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 50, 10, 52, 32, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 51, 32, 13, 10, 53, 32, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 52, 32, 105, 197, 159, 195, 182, 196, 159},
		},
		{
			name:           "-s flag",
			flags:          []string{execName, "cat", testFileName, "-s"},
			expectedOutput: []byte{116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 49, 32, 10, 32, 9, 116, 101, 115, 116, 32, 116, 97, 98, 10, 32, 10, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 50, 10, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 51, 32, 13, 10, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 52, 32, 105, 197, 159, 195, 182, 196, 159},
		},
		{
			name:           "-t flag",
			flags:          []string{execName, "cat", testFileName, "-t"},
			expectedOutput: []byte{116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 49, 32, 10, 32, 94, 73, 116, 101, 115, 116, 32, 116, 97, 98, 10, 32, 10, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 50, 10, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 51, 32, 94, 77, 10, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 52, 32, 105, 77, 45, 69, 77, 45, 94, 95, 77, 45, 67, 77, 45, 54, 77, 45, 68, 77, 45, 94, 95},
		},
		{
			name:           "-T flag",
			flags:          []string{execName, "cat", testFileName, "-T"},
			expectedOutput: []byte{116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 49, 32, 10, 32, 94, 73, 116, 101, 115, 116, 32, 116, 97, 98, 10, 32, 10, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 50, 10, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 51, 32, 13, 10, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 52, 32, 105, 197, 159, 195, 182, 196, 159},
		},
		{
			name:           "-v flag",
			flags:          []string{execName, "cat", testFileName, "-v"},
			expectedOutput: []byte{116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 49, 32, 10, 32, 9, 116, 101, 115, 116, 32, 116, 97, 98, 10, 32, 10, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 50, 10, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 51, 32, 94, 77, 10, 32, 116, 101, 115, 116, 32, 108, 105, 110, 101, 32, 52, 32, 105, 77, 45, 69, 77, 45, 94, 95, 77, 45, 67, 77, 45, 54, 77, 45, 68, 77, 45, 94, 95},
		},
	}

	for _, test := range testCases {
		fname := filepath.Join(os.TempDir(), "stdout")
		old := os.Stdout
		temp, _ := os.Create(fname)
		os.Stdout = temp

		err = app.Run(test.flags)

		temp.Close()
		os.Stdout = old
		out, _ := ioutil.ReadFile(fname)
		assert.Equal(t, test.expectedErr, err)
		assert.Equal(t, test.expectedOutput, out)
	}

}

func createTestFile(name, content string) error {
	return os.WriteFile(name, []byte(content), 0755)
}

func removeTestFile(name string) error {
	return os.Remove(name)

}
