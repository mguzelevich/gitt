package merge

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/mguzelevich/gitt/parser"
)

const (
	GSMERGE_UP_TO_DATE parser.SectionName = "GSMERGE_UP_TO_DATE"
)

func initGitStatusDscr(dscr *parser.Descriptor) error {
	dscr.AsString = gitStatusAsString
	return nil
}

func initGitStatusParser(p *parser.OutputParser) error {

	p.RegSection(GSMERGE_UP_TO_DATE,
		[]parser.OutLineRE{
			parser.NewRE("", `^Everything up-to-date$`),
		}, nil,
		func(sectionName parser.SectionName, name string, matches []string, dscr *parser.Descriptor) error {
			// dscr.SetString(GPULL_REPO_NAME, matches[1])
			return nil
		},
		parser.DummyHandler)
	return nil
}

func gitStatusAsString(dscr *parser.Descriptor, useAnsiColors bool) string {
	output := bytes.Buffer{}
	fmt.Fprintf(&output, "[OK]")
	return output.String()
}

func printMatches(name string, matches []string) {
	fmt.Fprintf(os.Stderr, "> %s : %s\n", name, strings.Join(matches, "|"))
}

func NewParser() *parser.OutputParser {
	return parser.NewParser(initGitStatusParser, initGitStatusDscr)
}
