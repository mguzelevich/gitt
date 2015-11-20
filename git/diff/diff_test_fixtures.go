package diff

type gitOutputTest struct {
	input  string
	output string
}

var testGitDiffFixtures = []gitOutputTest{
	gitOutputTest{`On branch BE-242`, `[BE-242] == []`},
}
