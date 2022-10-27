package cat

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

const (
	CmdEcho = "cat"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:        CmdEcho,
		HelpName:    CmdEcho,
		Action:      Action,
		ArgsUsage:   ` `,
		Usage:       `concatenate files and print on the standard output`,
		Description: `concatenate files and print on the standard output`,
		Flags:       Flags(),
	}
}

func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:     "A",
			Usage:    "equivalent to -vET",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "b",
			Usage:    "number nonempty output lines, overrides -n",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "e",
			Usage:    "equivalent to -vE",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "E",
			Usage:    "display $ at end of each line",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "n",
			Usage:    "number all output lines",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "s",
			Usage:    "suppress repeated empty output lines",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "t",
			Usage:    "equivalent to -vT",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "T",
			Usage:    "display TAB characters as ^I",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "v",
			Usage:    "use ^ and M- notation, except for LFD and TAB",
			Required: false,
		},
	}
}

func Action(c *cli.Context) error {
	args := c.Args()
	argsLen := args.Len()
	filePath := args.First()

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return errors.New("no such file")
	}

	fileByte, err := os.ReadFile(filePath)
	file := string(fileByte)

	if err != nil {
		return err
	}

	if argsLen == 1 {
		fmt.Print(file)
		return nil
	}

	argsMap := strSliceToMap(args.Tail())

	lines := make([][]byte, 0)
	line := make([]byte, 0)

	for _, b := range fileByte {
		line = append(line, b)
		if b == '\n' {
			lines = append(lines, line)
			line = nil
		}
	}

	if line != nil {
		lines = append(lines, line)
	}

	lineCounter := 1

	oldLine := make([]byte, 0)
	for _, line := range lines {
		lineStr := string(line)

		if argsMap["-s"] && (strings.TrimSpace(lineStr) == "" && strings.TrimSpace(string(oldLine)) == "") {
			continue
		}

		if argsMap["-n"] && !argsMap["-b"] {
			line = addNumberForLine(line, &lineCounter)
		}

		if argsMap["-b"] {
			line = addNumberForNonEmptyLine(line, &lineCounter)
		}

		if argsMap["-T"] {
			line = displayTabCharacter(line)
		}

		if argsMap["-E"] {
			line = displayEndOfLine(line)
		}

		if argsMap["-A"] {
			line = displayEndOfLine(line)
			line = displayTabCharacter(line)
			line = displayNoNPrt(line, true)

		}

		if argsMap["-e"] {
			line = displayEndOfLine(line)
			line = displayNoNPrt(line, argsMap["-T"])
		}

		if argsMap["-t"] {
			line = displayTabCharacter(line)
			line = displayNoNPrt(line, true)
		}

		if argsMap["-v"] {
			line = displayNoNPrt(line, argsMap["-T"])
		}

		fmt.Print(string(line))
		oldLine = line
		lineCounter++
	}

	return nil
}

func addNumberForLine(line []byte, counter *int) []byte {
	byteCounter := []byte(strconv.Itoa(*counter))
	byteCounter = append(byteCounter, ' ')
	return append(byteCounter, line...)
}

func addNumberForNonEmptyLine(line []byte, counter *int) []byte {
	if line == nil || strings.TrimSpace(string(line)) == "" {
		*counter--
		return line
	}

	return addNumberForLine(line, counter)
}

func displayTabCharacter(line []byte) []byte {

	for i := 0; i < len(line); i++ {
		b := line[i]
		if b == '\t' {
			line = remove(line, i)
			line = insert(line, i, '^')
			line = insert(line, i+1, 'I')
			i--
		}
	}

	return line
}

func displayEndOfLine(line []byte) []byte {
	lenOfLine := len(line)

	if lenOfLine == 1 && line[0] == '\n' {
		return insert(line, 0, '$')
	} else if lenOfLine == 1 && line[0] != '\n' {
		return line
	}

	lastChar, penultimateChar := line[lenOfLine-1], line[lenOfLine-2]
	if lastChar == '\n' && penultimateChar == '\r' {
		line = insert(line, lenOfLine-2, '^')
		line = insert(line, lenOfLine-2, 'M')
		return insert(line, lenOfLine-2, '$')

	}

	if lastChar == '\n' {
		return insert(line, lenOfLine-1, '$')
	}

	return line
}

func displayNoNPrt(line []byte, showTabs bool) []byte {
	temp := make([]byte, 0)

	for _, ch := range line {
		if ch >= 32 {
			if ch < 127 {
				temp = append(temp, ch)
			} else if ch == 127 {
				temp = append(temp, '^')
				temp = append(temp, '?')
			} else {
				temp = append(temp, 'M')
				temp = append(temp, '-')
				if ch >= 128+32 {
					if ch < 128+127 {
						temp = append(temp, ch-128)
					} else {
						temp = append(temp, '^')
						temp = append(temp, '?')
					}
				} else {
					temp = append(temp, '^')
					temp = append(temp, ch-128+64)
				}
			}
		} else if ch == '\t' && !showTabs {
			temp = append(temp, '\t')
		} else if ch == '\n' {
			temp = append(temp, '\n')
		} else {
			temp = append(temp, '^')
			temp = append(temp, ch+64)
		}

	}

	return temp
}

func strSliceToMap(args []string) map[string]bool {
	argsMap := make(map[string]bool)

	for _, a := range args {
		argsMap[a] = true
	}

	return argsMap
}

func insert(slice []byte, index int, value byte) []byte {
	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = value
	return slice
}

func remove(slice []byte, index int) []byte {
	return append(slice[:index], slice[index+1:]...)
}
