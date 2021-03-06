package split

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/IgaguriMK/sentence-splitter/subcmd"
)

func init() {
	subcmd.AddSubCommand(new(Split))
}

type Split struct {
	inFile  *os.File
	outFile *os.File
}

func (_ *Split) Cmd() string {
	return "split"
}

func (_ *Split) Help() string {
	return "Split text to sentences."
}

func (sp *Split) Register(cc *kingpin.CmdClause) {
	cc.Flag("input", "Input file.").Short('i').FileVar(&sp.inFile)
	cc.Flag("output", "Output file.").Short('o').OpenFileVar(&sp.outFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
}

func (sp *Split) Run() {
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

		for len(line) > 0 {
			sentence, left, found := getSentence(line, '.')
			fmt.Fprintf(sp.outFile, "- %s\n", sentence)

			if found {
				fmt.Fprintln(sp.outFile, "+ ")
			}

			line = strings.TrimSpace(left)
		}
	}

}

func getSentence(text string, seps ...rune) (sentence string, left string, found bool) {
	runes := []rune(text)

	sentenceRunes := make([]rune, 0)

	for len(runes) > 0 {
		r := runes[0]
		runes = runes[1:]

		sentenceRunes = append(sentenceRunes, r)

		for _, sep := range seps {
			if r == sep {
				return string(sentenceRunes), string(runes), true
			}
		}
	}

	return string(sentenceRunes), "", false
}
