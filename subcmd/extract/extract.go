package extract

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/IgaguriMK/sentence-splitter/subcmd"
)

func init() {
	subcmd.AddSubCommand(new(Extract))
}

type Extract struct {
	inFile  *os.File
	outFile *os.File
}

func (_ *Extract) Cmd() string {
	return "extract"
}

func (_ *Extract) Help() string {
	return "Extract translated sentences."
}

func (sp *Extract) Register(cc *kingpin.CmdClause) {
	cc.Flag("input", "Input file.").Short('i').FileVar(&sp.inFile)
	cc.Flag("output", "Output file.").Short('o').OpenFileVar(&sp.outFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
}

func (sp *Extract) Run() {
	if sp.inFile == nil {
		sp.inFile = os.Stdin
	}
	defer sp.inFile.Close()

	if sp.outFile == nil {
		sp.outFile = os.Stdout
	}
	defer sp.outFile.Close()

	sc := bufio.NewScanner(sp.inFile)

	for sc.Scan() {
		line := sc.Text()

		if line == "" {
			fmt.Fprintln(sp.outFile)
			continue
		}

		if strings.HasPrefix(line, "+ ") {
			fmt.Fprintln(sp.outFile, strings.TrimPrefix(line, "+ "))
		}
	}

}
