package pull

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/mguzelevich/gitt/git"
	"github.com/mguzelevich/gitt/parser"
)

const (
	GPULL_OBJECTS                     parser.FieldName = "GPULL_OBJECTS"
	GPULL_COMPRESSING_OBJECTS_PERCENT parser.FieldName = "GPULL_COMPRESSING_OBJECTS_PERCENT"
	GPULL_COMPRESSING_OBJECTS_ITEMS   parser.FieldName = "GPULL_COMPRESSING_OBJECTS_ITEMS"
	GPULL_COMPRESSING_OBJECTS_TOTAL   parser.FieldName = "GPULL_COMPRESSING_OBJECTS_TOTAL"
	GPULL_TOTAL                       parser.FieldName = "GPULL_TOTAL"
	GPULL_TOTAL_DELTA                 parser.FieldName = "GPULL_TOTAL_DELTA"
	GPULL_REUSED                      parser.FieldName = "GPULL_REUSED"
	GPULL_REUSED_DELTA                parser.FieldName = "GPULL_REUSED_DELTA"
	GPULL_REPO_NAME                   parser.FieldName = "GPULL_REPO_NAME"
	GPULL_LOCAL_COMMIT                parser.FieldName = "GPULL_LOCAL_COMMIT"
	GPULL_REPO_EXISTED_BRANCHES       parser.FieldName = "GPULL_REPO_EXISTED_BRANCHES"
	GPULL_REPO_NEW_BRANCHES           parser.FieldName = "GPULL_REPO_NEW_BRANCHES"

	GPULL_FILES_TMP            parser.FieldName = "GPULL_FILES_TMP"
	GPULL_FF_SUMMARY_CHANGED   parser.FieldName = "GPULL_FF_SUMMARY_CHANGED"
	GPULL_FF_SUMMARY_DELETION  parser.FieldName = "GPULL_FF_SUMMARY_DELETION"
	GPULL_FF_SUMMARY_INSERTION parser.FieldName = "GPULL_FF_SUMMARY_INSERTION"

	GPULL_UP_TO_DATE        parser.FieldName = "GPULL_UP_TO_DATE"
	GPULL_UNPACKING_PERCENT parser.FieldName = "GPULL_UNPACKING_PERCENT"
	GPULL_UNPACKING_ITEMS   parser.FieldName = "GPULL_UNPACKING_ITEMS"
	GPULL_UNPACKING_TOTAL   parser.FieldName = "GPULL_UNPACKING_TOTAL"
)

const (
	GSPULL_REMOTE             parser.SectionName = "GSPULL_COUNTING_OBJECTS"
	GSPULL_UNTRACKING_OBJECTS parser.SectionName = "GSPULL_UNTRACKING_OBJECTS"
	GSPULL_REPO               parser.SectionName = "GSPULL_REPO"
	GSPULL_SUBMODULES         parser.SectionName = "GSPULL_SUBMODULES"
	GSPULL_UPDATING           parser.SectionName = "GSPULL_UPDATING"
	GSPULL_FAST_FORWARD       parser.SectionName = "GSPULL_FAST_FORWARD"
	GSPULL_ALREADY_UP_TO_DATE parser.SectionName = "GSPULL_ALREADY_UP_TO_DATE"
)

func separateFiles(data map[string]interface{}) (git.Files, git.Files, git.Files, git.Files) {
	created := git.Files{}
	renamed := git.Files{}
	modified := git.Files{}
	deleted := git.Files{}

	for _, fdi := range data {
		fd := fdi.(git.FileData)
		switch fd.Operation {
		default:
			fd.Operation = "*"
			modified = append(modified, fd)
		case "created":
			fd.Operation = "+"
			created = append(created, fd)
		case "modified":
			fd.Operation = "*"
			modified = append(modified, fd)
		case "renamed":
			fd.Operation = ">"
			renamed = append(renamed, fd)
		case "deleted":
			fd.Operation = "-"
			deleted = append(deleted, fd)
		}
	}

	sort.Sort(created)
	sort.Sort(renamed)
	sort.Sort(modified)
	sort.Sort(deleted)

	return created, renamed, modified, deleted

}

func initGitPullDscr(dscr *parser.Descriptor) error {
	dscr.AsString = gitPullAsString

	dscr.AddField(GPULL_REPO_NAME, "")
	dscr.AddField(GPULL_LOCAL_COMMIT, "")
	dscr.AddField(GPULL_UP_TO_DATE, "")

	dscr.AddField(GPULL_FF_SUMMARY_CHANGED, 0)
	dscr.AddField(GPULL_FF_SUMMARY_INSERTION, 0)
	dscr.AddField(GPULL_FF_SUMMARY_DELETION, 0)

	dscr.AddField(GPULL_OBJECTS, "")
	dscr.AddField(GPULL_COMPRESSING_OBJECTS_PERCENT, "")
	dscr.AddField(GPULL_COMPRESSING_OBJECTS_ITEMS, "")
	dscr.AddField(GPULL_COMPRESSING_OBJECTS_TOTAL, "")
	dscr.AddField(GPULL_TOTAL, "")
	dscr.AddField(GPULL_TOTAL_DELTA, "")
	dscr.AddField(GPULL_REUSED, "")
	dscr.AddField(GPULL_REUSED_DELTA, "")
	dscr.AddField(GPULL_REPO_EXISTED_BRANCHES, []interface{}{})
	dscr.AddField(GPULL_REPO_NEW_BRANCHES, []interface{}{})

	dscr.AddField(GPULL_FILES_TMP, map[string]interface{}{})

	dscr.AddField(GPULL_UNPACKING_PERCENT, "")
	dscr.AddField(GPULL_UNPACKING_ITEMS, "")
	dscr.AddField(GPULL_UNPACKING_TOTAL, "")

	return nil
}

func initGitPullParser(p *parser.OutputParser) error {
	p.SectionBodyRegexp = regexp.MustCompile(` .+$`)

	p.RegSection(GSPULL_REMOTE,
		[]parser.OutLineRE{
			parser.NewRE("counting", `^remote\: Counting objects\: ([0-9]+), done\.$`),
			parser.NewRE("compressing", `^remote: Compressing objects: ([0-9]+)% \((.+)/(.+)\), done\.$`),
			parser.NewRE("total", `^remote: Total ([0-9]+) \(delta ([0-9]+)\), reused ([0-9]+) \(delta ([0-9]+)\)$`),
		}, nil,
		func(sectionName parser.SectionName, name string, matches []string, dscr *parser.Descriptor) error {
			//printMatches(matches)
			switch name {
			case "counting":
				c, _ := strconv.Atoi(matches[1])
				dscr.SetField(GPULL_OBJECTS, c)
			case "compressing":
				c, _ := strconv.Atoi(matches[1])
				dscr.SetField(GPULL_COMPRESSING_OBJECTS_PERCENT, c)
				c, _ = strconv.Atoi(matches[2])
				dscr.SetField(GPULL_COMPRESSING_OBJECTS_ITEMS, c)
				c, _ = strconv.Atoi(matches[3])
				dscr.SetField(GPULL_COMPRESSING_OBJECTS_TOTAL, c)
			case "total":
				c, _ := strconv.Atoi(matches[1])
				dscr.SetField(GPULL_TOTAL, c)
				c, _ = strconv.Atoi(matches[2])
				dscr.SetField(GPULL_TOTAL_DELTA, c)
				c, _ = strconv.Atoi(matches[3])
				dscr.SetField(GPULL_REUSED, c)
				c, _ = strconv.Atoi(matches[4])
				dscr.SetField(GPULL_REUSED_DELTA, c)
			default:
				p.Failed = true
			}
			return nil

		}, parser.DummyHandler)

	p.RegSection(GSPULL_UNTRACKING_OBJECTS,
		[]parser.OutLineRE{
			parser.NewRE("", `^Unpacking objects: ([0-9]+)% \((.+)/(.+)\), done.$`),
		}, nil,
		func(sectionName parser.SectionName, name string, matches []string, dscr *parser.Descriptor) error {
			c, _ := strconv.Atoi(matches[1])
			dscr.SetField(GPULL_UNPACKING_PERCENT, c)
			c, _ = strconv.Atoi(matches[2])
			dscr.SetField(GPULL_UNPACKING_ITEMS, c)
			c, _ = strconv.Atoi(matches[3])
			dscr.SetField(GPULL_UNPACKING_TOTAL, c)
			return nil
		}, parser.DummyHandler)
	p.RegSection(GSPULL_REPO,
		[]parser.OutLineRE{
			parser.NewRE("", `^From (.+)$`),
		},
		[]parser.OutLineRE{
			parser.NewRE("existed", `^([0-9a-z]+\.+[0-9a-z]+) +(.+) +-> +(.+)$`),
			parser.NewRE("forced", `\+ ([0-9a-z]+\.+[0-9a-z]+) +(.+) +-> +(.+) +\(forced update\)$`),
			parser.NewRE("new_branch", `^\* \[new branch\] +(.+) +-> +(.+)$`),
			parser.NewRE("new_tag", `^\* \[new tag\] +(.+) +-> +(.+)$`),
		},
		func(sectionName parser.SectionName, name string, matches []string, dscr *parser.Descriptor) error {
			dscr.SetString(GPULL_REPO_NAME, matches[1])
			return nil
		},
		func(name string, matches []string, dscr *parser.Descriptor) error {
			switch name {
			case "existed":
				dscr.AppendItem(GPULL_REPO_EXISTED_BRANCHES, git.BranchesLink{matches[1], matches[2], matches[3]})
			case "forced":
				dscr.AppendItem(GPULL_REPO_NEW_BRANCHES, git.BranchesLink{"", matches[1], matches[2]})
			case "new_branch":
				dscr.AppendItem(GPULL_REPO_NEW_BRANCHES, git.BranchesLink{"", matches[1], matches[2]})
			case "new_tag":
				dscr.AppendItem(GPULL_REPO_NEW_BRANCHES, git.BranchesLink{"", matches[1], matches[2]})
			default:
				p.Failed = true
			}
			return nil
		})

	p.RegSection(GSPULL_SUBMODULES,
		[]parser.OutLineRE{
			parser.NewRE("", `^Fetching submodule (.+)$`),
			parser.NewRE("", `^From (.+)$`),
		},
		[]parser.OutLineRE{
			parser.NewRE("existed", `^([0-9a-z]+\.+[0-9a-z]+) +(.+) +-> +(.+)$`),
			parser.NewRE("forced", `\+ ([0-9a-z]+\.+[0-9a-z]+) +(.+) +-> +(.+) +\(forced update\)$`),
			parser.NewRE("new_branch", `^\* \[new branch\] +(.+) +-> +(.+)$`),
			parser.NewRE("new_tag", `^\* \[new tag\] +(.+) +-> +(.+)$`),
		},
		func(sectionName parser.SectionName, name string, matches []string, dscr *parser.Descriptor) error {
			// dscr.SetString(GPULL_REPO_NAME, matches[1])
			return nil
		},
		func(name string, matches []string, dscr *parser.Descriptor) error {
			// switch name {
			// case "existed":
			// 	dscr.AppendItem(GPULL_REPO_EXISTED_BRANCHES, git.BranchesLink{matches[1], matches[2], matches[3]})
			// case "forced":
			// 	dscr.AppendItem(GPULL_REPO_NEW_BRANCHES, git.BranchesLink{"", matches[1], matches[2]})
			// case "new_branch":
			// 	dscr.AppendItem(GPULL_REPO_NEW_BRANCHES, git.BranchesLink{"", matches[1], matches[2]})
			// case "new_tag":
			// 	dscr.AppendItem(GPULL_REPO_NEW_BRANCHES, git.BranchesLink{"", matches[1], matches[2]})
			// default:
			// 	p.Failed = true
			// }
			return nil
		})

	p.RegSection(GSPULL_UPDATING,
		[]parser.OutLineRE{
			parser.NewRE("", `^Updating (.+)$`),
		}, nil,
		func(sectionName parser.SectionName, name string, matches []string, dscr *parser.Descriptor) error {
			dscr.SetString(GPULL_LOCAL_COMMIT, matches[1])
			return nil
		},
		parser.DummyHandler)

	p.RegSection(GSPULL_FAST_FORWARD,
		[]parser.OutLineRE{
			parser.NewRE("", `^Fast-forward$`),
		},
		[]parser.OutLineRE{
			parser.NewRE("changes_rename", `^(.+) => (.+) +\| +([0-9]+) ([\+\-]+)$`),
			parser.NewRE("changes_bin", `^(.+) +\| +Bin ([0-9]+) -> ([0-9]+) byte[s]*$`),
			parser.NewRE("changes", `^(.+) +\| +([0-9]+)( [\+\-]+)*$`),
			parser.NewRE("summary", `^([0-9]+) file[s]* changed(, ([0-9]+) ((insertion[s]*)|(deletion[s]*))\([\+\-]\))(, ([0-9]+) ((insertion[s]*)|(deletion[s]*))\([\+\-]\))*$`),
			parser.NewRE("create", `^create mode ([0-9]+) (.+)`),
			parser.NewRE("rename_dir", `^rename (.+\/)*{(.+) => (.+)}(\/.+) \(([0-9]+)%\)`),
			parser.NewRE("rename", `^rename (.+) => (.+) \(([0-9]+)%\)`),
			parser.NewRE("delete", `^delete mode ([0-9]+) (.+)`),
		}, parser.DummyTitleHandler,
		func(name string, matches []string, dscr *parser.Descriptor) error {
			// printMatches(name, matches)
			var fdi interface{}
			switch name {
			case "changes_rename":
				oldFilename := matches[1]
				filename := matches[2]
				lines, _ := strconv.Atoi(matches[3])
				added, deleted, err := git.ParseDiffLine(matches[4])
				if err != nil {
					printMatches(name, matches)
					p.Failed = true
					return err
				}
				fd := git.FileData{
					Operation:    "renamed",
					Filename:     filename,
					OldFilename:  oldFilename,
					Lines:        lines,
					LinesAdded:   added,
					LinesDeleted: deleted,
				}
				dscr.SetMapItem(GPULL_FILES_TMP, filename, fd)
			case "changes_bin":
				filename := matches[1]
				before, _ := strconv.Atoi(matches[2])
				after, _ := strconv.Atoi(matches[3])
				fd := git.FileData{
					Operation:    "",
					Filename:     filename,
					LinesAdded:   before,
					LinesDeleted: after,
				}
				dscr.SetMapItem(GPULL_FILES_TMP, filename, fd)
			case "changes":
				// printMatches(name, matches)
				filename := matches[1]
				// if fdi = dscr.GetMapItem(GPULL_FILES_TMP, filename); fdi != nil {
				// 	printMatches(name, matches)
				// 	p.Failed = true
				// 	return errors.New("file duplicate find")
				// }
				lines, _ := strconv.Atoi(matches[2])
				added, deleted, err := git.ParseDiffLine(matches[3])
				if err != nil {
					printMatches(name, matches)
					p.Failed = true
					return err
				}
				fd := git.FileData{
					Operation:    "",
					Filename:     filename,
					Lines:        lines,
					LinesAdded:   added,
					LinesDeleted: deleted,
				}
				dscr.SetMapItem(GPULL_FILES_TMP, filename, fd)
			case "summary":
				total, _ := strconv.Atoi(matches[1])
				insertions := 0
				deletions := 0
				switch matches[4] {
				case "insertions":
					changes, _ := strconv.Atoi(matches[3])
					insertions += changes
				case "deletions":
					changes, _ := strconv.Atoi(matches[3])
					deletions += changes
				}
				switch matches[9] {
				case "insertions":
					changes, _ := strconv.Atoi(matches[8])
					insertions += changes
				case "deletions":
					changes, _ := strconv.Atoi(matches[8])
					deletions += changes
				}
				dscr.SetField(GPULL_FF_SUMMARY_CHANGED, total)
				dscr.SetField(GPULL_FF_SUMMARY_INSERTION, insertions)
				dscr.SetField(GPULL_FF_SUMMARY_DELETION, deletions)
			case "create":
				filename := matches[2]
				mode, _ := strconv.Atoi(matches[1])
				if fdi = dscr.GetMapItem(GPULL_FILES_TMP, filename); fdi == nil {
					for k, tfdi := range dscr.GetField(GPULL_FILES_TMP).(map[string]interface{}) {
						if k[:3] == "..." && strings.HasSuffix(filename, k[3:]) {
							fdi = tfdi
							fd := fdi.(git.FileData)
							fd.Filename = filename
							dscr.SetMapItem(GPULL_FILES_TMP, k, fd)
							break
						}
					}
					if fdi == nil {
						printMatches(name, matches)
						p.Failed = true
						return errors.New("file not find")
					}
				}
				fd := fdi.(git.FileData)
				fd.Mode = mode
				fd.Operation = "created"
				dscr.SetMapItem(GPULL_FILES_TMP, filename, fd)
			case "rename_dir":
				oldFilename := matches[1] + matches[2] + matches[4]
				filename := matches[1] + matches[3] + matches[4]
				diffPercent, _ := strconv.Atoi(matches[5])
				fmt.Println("v", oldFilename)
				fmt.Println("v", filename)
				if fdi = dscr.GetMapItem(GPULL_FILES_TMP, filename); fdi == nil {
					fmt.Println("f", filename)
					printMatches(name, matches)
					p.Failed = true
					return errors.New("file not find")
				}
				fd := fdi.(git.FileData)
				if oldFilename != fd.OldFilename && fd.Operation != "renamed" {
					fmt.Println("f", filename)
					printMatches(name, matches)
					p.Failed = true
					return errors.New("file not find")
				}
				fd.DiffPercent = diffPercent
				dscr.SetMapItem(GPULL_FILES_TMP, filename, fd)
			case "rename":
				oldFilename := matches[1]
				filename := matches[2]
				diffPercent, _ := strconv.Atoi(matches[3])
				if fdi = dscr.GetMapItem(GPULL_FILES_TMP, filename); fdi == nil {
					printMatches(name, matches)
					p.Failed = true
					return errors.New("file not find")
				}
				fd := fdi.(git.FileData)
				if oldFilename != fd.OldFilename && fd.Operation != "renamed" {
					printMatches(name, matches)
					p.Failed = true
					return errors.New("file not find")
				}
				fd.DiffPercent = diffPercent
				dscr.SetMapItem(GPULL_FILES_TMP, filename, fd)
			case "delete":
				filename := matches[2]
				mode, _ := strconv.Atoi(matches[1])
				if fdi = dscr.GetMapItem(GPULL_FILES_TMP, filename); fdi == nil {
					printMatches(name, matches)
					p.Failed = true
					return errors.New("file not find")
				}
				fd := fdi.(git.FileData)
				fd.Mode = mode
				fd.Operation = "deleted"
				dscr.SetMapItem(GPULL_FILES_TMP, filename, fd)
			default:
				printMatches(name, matches)
				p.Failed = true
				return errors.New("GSPULL_FAST_FORWARD body parsing error")
			}
			return nil
		})
	p.RegSection(GSPULL_ALREADY_UP_TO_DATE,
		[]parser.OutLineRE{
			parser.NewRE("", `^Already up-to-date\.$`),
		}, nil,
		func(sectionName parser.SectionName, name string, matches []string, dscr *parser.Descriptor) error {
			dscr.SetString(GPULL_UP_TO_DATE, "true")
			return nil
		}, parser.DummyHandler)
	return nil
}

func gitPullAsString(dscr *parser.Descriptor, useAnsiColors bool) string {
	output := bytes.Buffer{}
	upToDate := dscr.GetString(GPULL_UP_TO_DATE)
	if upToDate != "" {
		dscr.AddToOut(&output, "up-to-date\n", parser.AnsiColor("green", useAnsiColors))
		return output.String()
	}

	repo := dscr.GetString(GPULL_REPO_NAME)
	localCommit := dscr.GetString(GPULL_LOCAL_COMMIT)

	fmt.Fprintf(&output, "[%s (%s)]", repo, localCommit)

	files := dscr.GetInt(GPULL_FF_SUMMARY_CHANGED)
	insertions := dscr.GetInt(GPULL_FF_SUMMARY_INSERTION)
	deletions := dscr.GetInt(GPULL_FF_SUMMARY_DELETION)
	fmt.Fprintf(&output, " %d*=%d+/%d-\n", files, insertions, deletions)

	// branches
	for _, bi := range dscr.GetField(GPULL_REPO_EXISTED_BRANCHES).([]interface{}) {
		b := bi.(git.BranchesLink)
		fmt.Fprintf(&output, "\t%s %s -> %s\n", b.Hash, b.Local, b.Remote)
	}
	for _, bi := range dscr.GetField(GPULL_REPO_NEW_BRANCHES).([]interface{}) {
		b := bi.(git.BranchesLink)
		fmt.Fprintf(&output, "\t+ [new] %s -> %s\n", b.Local, b.Remote)
	}

	// files
	created, renamed, modified, deleted := separateFiles(dscr.GetField(GPULL_FILES_TMP).(map[string]interface{}))
	if len(created) > 0 || len(renamed) > 0 || len(modified) > 0 || len(deleted) > 0 {
		fmt.Fprintf(&output, "changes:\n")
	}
	for _, fd := range created {
		fmt.Fprintf(&output, "\t%s %s %d*=%d+/%d-", fd.Operation, fd.Filename, fd.Lines, fd.LinesAdded, fd.LinesDeleted)
		if fd.Mode != 0 {
			fmt.Fprintf(&output, " (%d)", fd.Mode)
		}
		fmt.Fprintf(&output, "\n")
	}
	for _, fd := range renamed {
		fmt.Fprintf(&output, "\t%s %s %d*=%d+/%d-", fd.Operation, fd.Filename, fd.Lines, fd.LinesAdded, fd.LinesDeleted)
		if fd.Mode != 0 {
			fmt.Fprintf(&output, " (%d)", fd.Mode)
		}
		fmt.Fprintf(&output, "\n")
	}
	for _, fd := range modified {
		fmt.Fprintf(&output, "\t%s %s %d*=%d+/%d-", fd.Operation, fd.Filename, fd.Lines, fd.LinesAdded, fd.LinesDeleted)
		if fd.Mode != 0 {
			fmt.Fprintf(&output, " (%d)", fd.Mode)
		}
		fmt.Fprintf(&output, "\n")
	}
	for _, fd := range deleted {
		fmt.Fprintf(&output, "\t%s %s %d*=%d+/%d-", fd.Operation, fd.Filename, fd.Lines, fd.LinesAdded, fd.LinesDeleted)
		if fd.Mode != 0 {
			fmt.Fprintf(&output, " (%d)", fd.Mode)
		}
		fmt.Fprintf(&output, "\n")
	}
	return output.String()
}

func printMatches(name string, matches []string) {
	fmt.Fprintf(os.Stderr, "> %s : %s\n", name, strings.Join(matches, "|"))
}

func NewParser() *parser.OutputParser {
	return parser.NewParser(initGitPullParser, initGitPullDscr)
}
