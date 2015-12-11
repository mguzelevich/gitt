package status

const (
	TMPL_STATUS = `
= {{.Idx}} = {{.Path}} [{{.LocalBranch}}] == [{{.RemoteBranch}}]
{{with .ToCommit}}
  changes to be committed:
  	{{range $item := .ToCommit.FilesAdd}}
    + {{$item}}
  	{{end}}
  	{{range $item := .ToCommit.FilesModify}}
    * {{$item}}
  	{{end}}
  	{{range $item := .ToCommit.FilesRename}}
    > {{$item}}
  	{{end}}
  	{{range $item := .ToCommit.FilesDelete}}
    - {{$item}}
  	{{end}}
{{end}}
{{with .NotStaged}}
  changes not staged for commit:  
  	{{range $item := .NotStaged.FilesAdd}}
    + {{$item}}
  	{{end}}
  	{{range $item := .NotStaged.FilesModify}}
    * {{$item}}
  	{{end}}
  	{{range $item := .NotStaged.FilesRename}}
    > {{$item}}
  	{{end}}
  	{{range $item := .NotStaged.FilesDelete}}
    - {{$item}}
  	{{end}}
{{end}}
{{with .Untracked}}
  untracked files:
  	{{range $item := .Untracked.Files}}
    ? {{$item}}
  	{{end}}
{{end}}`

	TMPL_STATUS_COLOR = `

= %

`
)

/*
func func_name() {
	fmt.Fprintf(&output, "[")

	branch := dscr.GetString(GST_BRANCH)
	if branch != "master" {
		dscr.AddToOut(&output, branch, parser.AnsiColor("yellow", useAnsiColors))
	} else {
		fmt.Fprintf(&output, "%s", branch)
	}
	pairBranch := dscr.GetString(GST_PAIR_BRANCH)
	aheadCommits := dscr.GetInt(GST_AHEAD_COMMITS)
	remoteCommits := dscr.GetInt(GST_REMOTE_COMMITS)
	if aheadCommits > 0 {
		dscr.AddToOut(&output, fmt.Sprintf("+%d", aheadCommits), parser.AnsiColor("green", useAnsiColors))
	}
	if aheadCommits == 0 && remoteCommits == 0 {
		fmt.Fprintf(&output, "] == [%s", pairBranch)
	} else {
		fmt.Fprintf(&output, "] ~~ [%s", pairBranch)
	}
	if remoteCommits > 0 {
		dscr.AddToOut(&output, fmt.Sprintf("+%d", remoteCommits), parser.AnsiColor("green", useAnsiColors))
	}
	fmt.Fprintf(&output, "]\n")

	toCommitNew := dscr.GetField(GST_TO_COMMIT_NEW).([]interface{})
	toCommitModified := dscr.GetField(GST_TO_COMMIT_MODIFIED).([]interface{})
	toCommitRenamed := dscr.GetField(GST_TO_COMMIT_RENAMED).([]interface{})
	toCommitDeleted := dscr.GetField(GST_TO_COMMIT_DELETED).([]interface{})

	if len(toCommitNew) > 0 || len(toCommitModified) > 0 || len(toCommitRenamed) > 0 || len(toCommitDeleted) > 0 {
		fmt.Fprintf(&output, "  changes to be committed:\n")
		dscr.AddToOut(&output, parser.MakeFilesList(toCommitNew, "    + "), parser.AnsiColor("green", useAnsiColors))
		dscr.AddToOut(&output, parser.MakeFilesList(toCommitModified, "    * "), parser.AnsiColor("yellow", useAnsiColors))
		for _, ri := range toCommitRenamed {
			r := ri.(git.FileData)
			dscr.AddToOut(&output, "    > ", parser.AnsiColor("yellow", useAnsiColors))
			dscr.AddToOut(&output, r.OldFilename, parser.AnsiColor("red", useAnsiColors))
			fmt.Fprintf(&output, " -> ")
			dscr.AddToOut(&output, r.Filename, parser.AnsiColor("green", useAnsiColors))
			fmt.Fprintf(&output, "\n")
		}
		dscr.AddToOut(&output, parser.MakeFilesList(toCommitDeleted, "    - "), parser.AnsiColor("red", useAnsiColors))
	}

	notStagedNew := dscr.GetField(GST_NOT_STAGED_NEW).([]interface{})
	notStagedModified := dscr.GetField(GST_NOT_STAGED_MODIFIED).([]interface{})
	notStagedRenamed := dscr.GetField(GST_NOT_STAGED_RENAMED).([]interface{})
	notStagedDeleted := dscr.GetField(GST_NOT_STAGED_DELETED).([]interface{})
	if len(notStagedNew) > 0 || len(notStagedModified) > 0 || len(notStagedRenamed) > 0 || len(notStagedDeleted) > 0 {
		fmt.Fprintf(&output, "  changes not staged for commit:\n")
		dscr.AddToOut(&output, parser.MakeFilesList(notStagedNew, "    + "), parser.AnsiColor("green", useAnsiColors))
		dscr.AddToOut(&output, parser.MakeFilesList(notStagedModified, "    * "), parser.AnsiColor("yellow", useAnsiColors))
		for _, ri := range notStagedRenamed {
			r := ri.(git.FileData)
			dscr.AddToOut(&output, "    > ", parser.AnsiColor("yellow", useAnsiColors))
			dscr.AddToOut(&output, r.OldFilename, parser.AnsiColor("red", useAnsiColors))
			fmt.Fprintf(&output, " -> ")
			dscr.AddToOut(&output, r.Filename, parser.AnsiColor("green", useAnsiColors))
			fmt.Fprintf(&output, "\n")
		}
		dscr.AddToOut(&output, parser.MakeFilesList(notStagedDeleted, "    - "), parser.AnsiColor("red", useAnsiColors))
	}

	untracked := dscr.GetField(GST_UNTRACKED).([]interface{})
	if len(untracked) > 0 {
		fmt.Fprintf(&output, "  untracked files:\n")
		dscr.AddToOut(&output, parser.MakeFilesList(untracked, "    ? "), parser.AnsiColor("yellow", useAnsiColors))
	}

}
*/
