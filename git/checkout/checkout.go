package checkout

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/mguzelevich/gitt/parser"
)

const (
	GST_BRANCH      parser.FieldName = "GST_BRANCH"      // string
	GST_PAIR_BRANCH parser.FieldName = "GST_PAIR_BRANCH" // string
)

const (
	GSCHEKOUT_HEAD parser.SectionName = "GSCHEKOUT_HEAD"
)

func initGitStatusDscr(dscr *parser.Descriptor) error {
	dscr.AsString = gitStatusAsString

	dscr.AddField(GST_BRANCH, "")
	dscr.AddField(GST_PAIR_BRANCH, "")

	return nil
}

func initGitStatusParser(p *parser.OutputParser) error {

	p.RegSection(GSCHEKOUT_HEAD,
		[]parser.OutLineRE{
			parser.NewRE("already", `^Already on '(.+)'$`),
			parser.NewRE("switched", `^Switched to branch '(.+)'$`),
			parser.NewRE("up-to-date", `^Your branch is up-to-date with '(.+)'.`),
		}, nil,
		func(sectionName parser.SectionName, name string, matches []string, dscr *parser.Descriptor) error {
			switch name {
			case "already":
				dscr.SetString(GST_BRANCH, matches[1])
			case "switched":
				dscr.SetString(GST_BRANCH, matches[1])
			case "up-to-date":
				dscr.SetString(GST_PAIR_BRANCH, matches[1])
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
	pairBranch := dscr.GetString(GST_PAIR_BRANCH)
	if branch != "master" {
		dscr.AddToOut(&output, branch, parser.AnsiColor("yellow", useAnsiColors))
	} else {
		fmt.Fprintf(&output, "%s", branch)
	}

	fmt.Fprintf(&output, "] == [%s", pairBranch)
	fmt.Fprintf(&output, "]\n")

	return output.String()
}

func printMatches(name string, matches []string) {
	fmt.Fprintf(os.Stderr, "> %s : %s\n", name, strings.Join(matches, "|"))
}

func NewParser() *parser.OutputParser {
	return parser.NewParser(initGitStatusParser, initGitStatusDscr)
}
