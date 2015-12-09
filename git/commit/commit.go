package commit

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/mguzelevich/gitt/parser"
)

const (
	GST_BRANCH         parser.FieldName = "GST_BRANCH"         // string
	GST_BRANCH_HASH    parser.FieldName = "GST_BRANCH_HASH"    // string
	GST_COMMIT_MESSAGE parser.FieldName = "GST_COMMIT_MESSAGE" // string
)

const (
	GSCOMMIT_HEAD parser.SectionName = "GSCOMMIT_HEAD"
)

func initGitCommitDscr(dscr *parser.Descriptor) error {
	dscr.AsString = gitStatusAsString

	dscr.AddField(GST_BRANCH, "")
	dscr.AddField(GST_BRANCH_HASH, "")
	dscr.AddField(GST_COMMIT_MESSAGE, "")

	return nil
}

func initGitCommitParser(p *parser.OutputParser) error {

	p.RegSection(GSCOMMIT_HEAD,
		[]parser.OutLineRE{
			parser.NewRE("commit", `^[[(.+) (.+)] (.+)$`),
		}, nil,
		func(sectionName parser.SectionName, name string, matches []string, dscr *parser.Descriptor) error {
			switch name {
			case "commit":
				dscr.SetString(GST_BRANCH, matches[1])
				dscr.SetString(GST_BRANCH_HASH, matches[2])
				dscr.SetString(GST_COMMIT_MESSAGE, matches[3])
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
	return parser.NewParser(initGitCommitParser, initGitCommitDscr)
}
