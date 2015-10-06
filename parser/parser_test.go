package parser

import (
	"bytes"
	"fmt"
	//	"os"
	"testing"
)

const (
	TST_FIELD FieldName = "TST_FIELD" // string
)

const (
	SECTION_A SectionName = "SECTION_A"
	SECTION_B SectionName = "SECTION_B"
	SECTION_C SectionName = "SECTION_C"
)

var testOutput = bytes.Buffer{}

func initDscr(dscr *Descriptor) error {
	dscr.AsString = asString

	dscr.AddField(TST_FIELD, "")

	return nil
}

func initParser(p *OutputParser) error {
	p.RegSection(SECTION_A,
		[]OutLineRE{
			NewRE("a", `^section A. title [0-9]$`),
			NewRE("b", `^section A. title [0-9][0-9]$`),
			NewRE("c", `^section A. title [0-9]+$`),
		}, nil,
		func(section SectionName, name string, matches []string, dscr *Descriptor) error {
			fmt.Fprintf(&testOutput, "t: %s %s %s\n", section, name, matches)
			return nil
		},
		func(name string, matches []string, dscr *Descriptor) error {
			fmt.Fprintf(&testOutput, "b: %s %s\n", name, matches)
			return nil

		})
	p.RegSection(SECTION_B,
		[]OutLineRE{
			NewRE("a", `^section B. title [0-9]$`),
		}, nil,
		func(sectionName SectionName, name string, matches []string, dscr *Descriptor) error {
			fmt.Fprintf(&testOutput, "t: %s %s %s\n", sectionName, name, matches)
			return nil

		},
		func(name string, matches []string, dscr *Descriptor) error {
			fmt.Fprintf(&testOutput, "b: %s %s\n", name, matches)
			return nil
		})
	p.RegSection(SECTION_C,
		[]OutLineRE{
			NewRE("a", `^section C. title (.*)$`),
		}, nil,
		func(sectionName SectionName, name string, matches []string, dscr *Descriptor) error {
			fmt.Fprintf(&testOutput, "t: %s %s %s\n", sectionName, name, matches)
			return nil

		},
		func(name string, matches []string, dscr *Descriptor) error {
			fmt.Fprintf(&testOutput, "b: %s %s\n", name, matches)
			return nil

		})
	return nil
}

func asString(dscr *Descriptor, useAnsiColors bool) string {
	output := bytes.Buffer{}
	fmt.Fprintf(&output, "[")

	branch := dscr.GetString(TST_FIELD)
	if branch != "master" {
		dscr.AddToOut(&output, branch, AnsiColor("yellow", useAnsiColors))
	} else {
		fmt.Fprintf(&output, "%s", branch)
	}

	return output.String()
}

func TestOutParser_status(t *testing.T) {
	for testIdx, fixture := range testFixtures {
		testOutput = bytes.Buffer{}

		buffer := bytes.Buffer{}
		buffer.Write([]byte(fixture.input))

		p := NewParser(initParser, initDscr)

		out, _ := p.splitOutput(buffer)
		if out.String() != fixture.output {
			t.Error(
				"For", testIdx,
				"expected", fixture.input,
				"got", out.String(),
				//"output", debugOutput.String(),
			)
		}
		// fmt.Fprintf(os.Stderr, "\n=== %d\n%s\n", testIdx, testOutput.String())
		// fmt.Fprintf(os.Stderr, "\n=== %d\n%s\n", testIdx, out.String())
	}
}
