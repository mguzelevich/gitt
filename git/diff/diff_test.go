package diff

import (
	"bytes"
	"fmt"
	//"os"
	"testing"

	"github.com/mguzelevich/gitt/parser"
)

func TestGitDiff_1(t *testing.T) {
	for i, fixture := range testGitDiffFixtures {
		buffer := bytes.Buffer{}
		buffer.Write([]byte(fixture.input))

		p := NewParser()
		dscr := parser.NewDescriptor()

		if dscr, err := p.Process(buffer, dscr); err != nil {
			t.Error(
				"For", i,
				"expected", fixture.output,
				"got", err,
				"output\n", p.DebugOutput.String(),
			)
		} else if p.Failed {
			t.Error(
				"For", i,
				"expected", fixture.output,
				"got", err,
				"output\n", p.DebugOutput.String(),
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
	}
}

// func TestMain(m *testing.M) {
// 	sp.Init()

// 	os.Exit(m.Run())
// }
