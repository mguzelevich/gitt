package branch

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/mguzelevich/gitt/parser"
)

const (
	GST_BRANCH parser.FieldName = "GST_BRANCH" // string
)

const (
	GSBRANCH_HEAD parser.SectionName = "GSBRANCH_HEAD"
)

func initGitBranchDscr(dscr *parser.Descriptor) error {
	dscr.AsString = gitStatusAsString

	dscr.AddField(GST_BRANCH, "")

	return nil
}

func initGitBranchParser(p *parser.OutputParser) error {

	p.RegSection(GSBRANCH_HEAD,
		[]parser.OutLineRE{
			parser.NewRE("deleted", `^Deleted branch (.+) \(was (.+)\).$`),
		}, nil,
		func(sectionName parser.SectionName, name string, matches []string, dscr *parser.Descriptor) error {
			switch name {
			case "deleted":
				dscr.SetString(GST_BRANCH, matches[1])
				//dscr.SetString(GST_BRANCH, matches[2])
			default:
				p.Failed = true
			}
			return nil
		},
		parser.DummyHandler)
	return nil
}

func gitStatusAsString(dscr *parser.Descriptor, useAnsiColors bool) string {
	output := bytes.Buffer{}
	fmt.Fprintf(&output, "[")

	branch := dscr.GetString(GST_BRANCH)
	if branch != "master" {
		dscr.AddToOut(&output, branch, parser.AnsiColor("yellow", useAnsiColors))
	} else {
		fmt.Fprintf(&output, "%s", branch)
	}

	fmt.Fprintf(&output, "]\n")

	return output.String()
}

func printMatches(name string, matches []string) {
	fmt.Fprintf(os.Stderr, "> %s : %s\n", name, strings.Join(matches, "|"))
}

func NewParser() *parser.OutputParser {
	return parser.NewParser(initGitBranchParser, initGitBranchDscr)
}
