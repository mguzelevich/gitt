package diff

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mguzelevich/gitt/git"
	"github.com/mguzelevich/gitt/parser"
)

const (
	GST_BRANCH              parser.FieldName = "GST_BRANCH"              // string
	GST_PAIR_BRANCH         parser.FieldName = "GST_PAIR_BRANCH"         // string
	GST_AHEAD_COMMITS       parser.FieldName = "GST_AHEAD_COMMITS"       // int
	GST_REMOTE_COMMITS      parser.FieldName = "GST_REMOTE_COMMITS"      // int
	GST_TO_COMMIT_NEW       parser.FieldName = "GST_TO_COMMIT_NEW"       // []string
	GST_TO_COMMIT_MODIFIED  parser.FieldName = "GST_TO_COMMIT_MODIFIED"  // []string
	GST_TO_COMMIT_RENAMED   parser.FieldName = "GST_TO_COMMIT_RENAMED"   // []renamedFile
	GST_TO_COMMIT_DELETED   parser.FieldName = "GST_TO_COMMIT_DELETED"   // []string
	GST_NOT_STAGED_NEW      parser.FieldName = "GST_NOT_STAGED_NEW"      // []string
	GST_NOT_STAGED_MODIFIED parser.FieldName = "GST_NOT_STAGED_MODIFIED" // []string
	GST_NOT_STAGED_RENAMED  parser.FieldName = "GST_NOT_STAGED_RENAMED"  // []renamedFile
	GST_NOT_STAGED_DELETED  parser.FieldName = "GST_NOT_STAGED_DELETED"  // []string
	GST_UNTRACKED           parser.FieldName = "GST_UNTRACKED"           // []string
)

const (
	GSST_BRANCH_NAME     parser.SectionName = "GSST_BRANCH_NAME"
	GSST_BRANCH_DIFF     parser.SectionName = "GSST_BRANCH_DIFF"
	GSST_BRANCH_DIVIRGED parser.SectionName = "GSST_BRANCH_DIVIRGED"
	GSST_TO_BE_COMMITED  parser.SectionName = "GSST_TO_BE_COMMITED"
	GSST_NOT_STAGED      parser.SectionName = "GSST_NOT_STAGED"
	GSST_UNTRACKED       parser.SectionName = "GSST_UNTRACKED"
	GSST_FOOTER          parser.SectionName = "GSST_FOOTER"
)

func initGitStatusDscr(dscr *parser.Descriptor) error {
	dscr.AsString = gitStatusAsString

	dscr.AddField(GST_BRANCH, "")
	dscr.AddField(GST_PAIR_BRANCH, "")
	dscr.AddField(GST_AHEAD_COMMITS, 0)
	dscr.AddField(GST_REMOTE_COMMITS, 0)
	dscr.AddField(GST_TO_COMMIT_NEW, []interface{}{})
	dscr.AddField(GST_TO_COMMIT_MODIFIED, []interface{}{})
	dscr.AddField(GST_TO_COMMIT_RENAMED, []interface{}{})
	dscr.AddField(GST_TO_COMMIT_DELETED, []interface{}{})
	dscr.AddField(GST_NOT_STAGED_NEW, []interface{}{})
	dscr.AddField(GST_NOT_STAGED_MODIFIED, []interface{}{})
	dscr.AddField(GST_NOT_STAGED_RENAMED, []interface{}{})
	dscr.AddField(GST_NOT_STAGED_DELETED, []interface{}{})
	dscr.AddField(GST_UNTRACKED, []interface{}{})

	return nil
}

func initGitStatusParser(p *parser.OutputParser) error {

	p.RegSection(GSST_BRANCH_NAME,
		[]parser.OutLineRE{
			parser.NewRE("", `^On branch (.*)$`),
		}, nil,
		func(sectionName parser.SectionName, name string, matches []string, dscr *parser.Descriptor) error {
			//printMatches(name, matches)
			dscr.SetString(GST_BRANCH, matches[1])
			return nil
		}, parser.DummyHandler)

	p.RegSection(GSST_BRANCH_DIFF,
		[]parser.OutLineRE{
			parser.NewRE("initial", `^Initial commit$`),
			parser.NewRE("ahead", `^Your branch is ahead of '(.+)'( by ([0-9]+) commit[s]*)*.$`),
			parser.NewRE("up_to_date", `^Your branch is up-to-date with '(.+)'\.$`),
		}, nil,
		func(sectionName parser.SectionName, name string, matches []string, dscr *parser.Descriptor) error {
			switch name {
			case "initial":
				dscr.SetString(GST_PAIR_BRANCH, "<initial commit>")
			case "ahead":
				dscr.SetString(GST_PAIR_BRANCH, matches[1])
				aheadCommits, _ := strconv.Atoi(matches[3])
				dscr.SetField(GST_AHEAD_COMMITS, aheadCommits)
			case "up_to_date":
				dscr.SetString(GST_PAIR_BRANCH, matches[1])
			default:
				p.Failed = true
			}
			/*
				switch {
				case matches[1] == "Initial commit":
					dscr.SetString(GST_PAIR_BRANCH, "<initial commit>")
				case matches[3] == "ahead of":
					dscr.SetString(GST_PAIR_BRANCH, matches[6])
					aheadCommits, _ := strconv.Atoi(matches[8])
					dscr.SetField(GST_AHEAD_COMMITS, aheadCommits)
				case matches[3] == "up-to-date with":
					dscr.SetString(GST_PAIR_BRANCH, matches[6])
				default:
					p.Failed = true
				}
			*/
			return nil
		}, parser.DummyHandler)
	p.RegSection(GSST_BRANCH_DIVIRGED,
		[]parser.OutLineRE{
			parser.NewRE("remote", `^Your branch and '(.*)' have diverged,$`),
			parser.NewRE("diff", `^and have (.*) and (.*) different commits each, respectively.$`),
		}, nil,
		func(sectionName parser.SectionName, name string, matches []string, dscr *parser.Descriptor) error {
			switch name {
			case "remote":
				dscr.SetString(GST_PAIR_BRANCH, matches[1])
			case "diff":
				//fmt.Println("d", strings.Join(matches, "|"))
				aheadCommits, _ := strconv.Atoi(matches[1])
				dscr.SetField(GST_AHEAD_COMMITS, aheadCommits)
				remoteCommits, _ := strconv.Atoi(matches[2])
				dscr.SetField(GST_REMOTE_COMMITS, remoteCommits)
			}
			return nil
		}, parser.DummyHandler)

	p.RegSection(GSST_TO_BE_COMMITED,
		[]parser.OutLineRE{
			parser.NewRE("", `^Changes to be committed:$`),
		},
		[]parser.OutLineRE{
			parser.NewRE("modified", `^modified: +(.*)$`),
			parser.NewRE("new", `^new file: +(.*)$`),
			parser.NewRE("deleted", `^deleted: +(.*)$`),
			parser.NewRE("renamed", `^renamed: +(.+) -> (.+)$`),
		},
		parser.DummyTitleHandler,
		func(name string, matches []string, dscr *parser.Descriptor) error {
			filename := strings.TrimSpace(matches[1]) // file
			switch name {
			case "new":
				dscr.AppendItem(GST_TO_COMMIT_NEW, filename)
			case "modified":
				dscr.AppendItem(GST_TO_COMMIT_MODIFIED, filename)
			case "renamed":
				dscr.AppendItem(GST_TO_COMMIT_RENAMED,
					git.FileData{
						OldFilename: matches[1],
						Filename:    matches[2],
						DiffPercent: 100,
					})
			case "deleted":
				dscr.AppendItem(GST_TO_COMMIT_DELETED, filename)
			default:
				p.Failed = true
			}
			return nil
		})

	p.RegSection(GSST_NOT_STAGED,
		[]parser.OutLineRE{
			parser.NewRE("", `^Changes not staged for commit:$`),
		},
		[]parser.OutLineRE{
			parser.NewRE("modified", `^modified: +(.*)$`),
			parser.NewRE("new", `^new file: +(.*)$`),
			parser.NewRE("deleted", `^deleted: +(.*)$`),
			parser.NewRE("renamed", `^renamed: +(.+) -> (.+)$`),
		},
		parser.DummyTitleHandler,
		func(name string, matches []string, dscr *parser.Descriptor) error {
			filename := matches[1] // file
			switch name {
			case "new":
				dscr.AppendItem(GST_NOT_STAGED_NEW, filename)
			case "modified":
				dscr.AppendItem(GST_NOT_STAGED_MODIFIED, filename)
			case "renamed":
				dscr.AppendItem(GST_NOT_STAGED_RENAMED,
					git.FileData{
						OldFilename: matches[1],
						Filename:    matches[2],
						DiffPercent: 100,
					})
			case "deleted":
				dscr.AppendItem(GST_NOT_STAGED_DELETED, filename)
			default:
				p.Failed = true
			}
			return nil
		})

	p.RegSection(GSST_UNTRACKED,
		[]parser.OutLineRE{
			parser.NewRE("", `^Untracked files:$`),
		},
		[]parser.OutLineRE{
			parser.NewRE("", `^(.*)$`),
		},
		parser.DummyTitleHandler,
		func(name string, matches []string, dscr *parser.Descriptor) error {
			filename := matches[1] // file
			dscr.AppendItem(GST_UNTRACKED, filename)
			return nil
		})

	p.RegSection(GSST_FOOTER,
		[]parser.OutLineRE{
			parser.NewRE("no_changes", `^no changes added to commit \(use "git add" and/or "git commit -a"\)$`),
			parser.NewRE("nothing_to_commit_clean", `^nothing to commit, working directory clean$`),
			parser.NewRE("nothing_to_commit", `^nothing to commit \(create/copy files and use "git add" to track\)$`),
			parser.NewRE("nothing_added_to_commit", `^nothing added to commit but untracked files present \(use "git add" to track\)$`),
		},
		nil, parser.DummyTitleHandler, parser.DummyHandler)
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
	return output.String()
}

func printMatches(name string, matches []string) {
	fmt.Fprintf(os.Stderr, "> %s : %s\n", name, strings.Join(matches, "|"))
}

func NewParser() *parser.OutputParser {
	return parser.NewParser(initGitStatusParser, initGitStatusDscr)
}
