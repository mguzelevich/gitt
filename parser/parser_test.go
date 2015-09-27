package parser

import (
	"bytes"
	"fmt"
	//"os"
	"testing"
)

const (
	TST_FIELD FieldName = "TST_FIELD" // string
)

const (
	SECTION_1 SectionName = "SECTION_1"
	SECTION_2 SectionName = "SECTION_2"
)

func initDscr(dscr *Descriptor) error {
	dscr.AsString = asString

	dscr.AddField(TST_FIELD, "")

	return nil
}

func initParser(p *OutputParser) error {
	p.RegSection(SECTION_1,
		[]OutLineRE{
			NewRE("", `^On branch (.*)$`),
		}, nil,
		func(name string, matches []string, dscr *Descriptor) error {
			dscr.SetString(TST_FIELD, matches[1])
			return nil
		}, DummyHandler)
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
		buffer := bytes.Buffer{}
		buffer.Write([]byte(fixture.input))

		p := NewParser(initParser, initDscr)

		out, _ := p.splitOutput(buffer)
		for i, section := range out {
			s, _ := p.identifySection(section.titles)
			fmt.Printf("[%03d] [%03d] [%s]\n", testIdx, i, s)
			fmt.Printf("titles:\n")
			for _, l := range section.titles {
				fmt.Printf("\t [%s]\n", l)
			}
			fmt.Printf("body:\n")
			for _, l := range section.body {
				fmt.Printf("\t [%s]\n", l)
			}
		}

		// t.Error(
		// 	"For", testIdx,
		// 	"got", out,
		// 	"output", debugOutput.String(),
		// )

		/*
			if dscr, err := Process(buffer, dscr); err != nil {
				t.Error(
					"For", i,
					"expected", fixture.output,
					"got", err,
					"output", debugOutput.String(),
				)
			} else if failed {
				t.Error(
					"For", i,
					"expected", fixture.output,
					"got", err,
					"output", debugOutput.String(),
				)
			} else {
				statusString := dscr.AsString(dscr, false)
				if statusString != fixture.output {
					// fmt.Printf("got [%s]\n", dscr.String())
					// fmt.Printf("exp [%s]\n", fixture.output)
					t.Error(
						fmt.Sprintf("For %d expected >%s< got >%s<", i, fixture.output, statusString),
					)
				}
			}
			//fmt.Println("\n")
		*/
	}
}
