package parser

type gitOutputTest struct {
	input  string
	output string
}

var testFixtures = []gitOutputTest{
	gitOutputTest{`section A. title 1
section A. title 12
section A. title 123
section B. title 1
section C. title 1
`, `[000] SECTION_A
t: a
t: b
t: c
[001] SECTION_B
t: a
[002] SECTION_C
t: a
`},
	gitOutputTest{`section A. title 1

section A. title 12
`, `[000] SECTION_A
t: a
[001] SECTION_A
t: b
`},
	gitOutputTest{`section A. title 1

section B. title 1

section C. title 1
`, `[000] SECTION_A
t: a
[001] SECTION_B
t: a
[002] SECTION_C
t: a
`},
}
