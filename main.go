package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mguzelevich/gitt/parser"

	"github.com/mguzelevich/gitt/git/branch"
	"github.com/mguzelevich/gitt/git/checkout"
	"github.com/mguzelevich/gitt/git/diff"
	"github.com/mguzelevich/gitt/git/pull"
	"github.com/mguzelevich/gitt/git/push"
	"github.com/mguzelevich/gitt/git/rebase"
	"github.com/mguzelevich/gitt/git/status"
)

type Action struct {
	Cmd  string
	Args []string
}

var GIT_CHECKOUT = Action{Cmd: "checkout"}
var GIT_DIFF = Action{Cmd: "diff"}
var GIT_PULL = Action{Cmd: "pull"}
var GIT_PUSH = Action{Cmd: "push"}
var GIT_REBASE = Action{Cmd: "rebase"}
var GIT_STATUS = Action{Cmd: "status"}
var GIT_BRANCH = Action{Cmd: "branch"}

var dirs = []string{}
var gitbinary string

func checkError(err error) {
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}

func walkerGetGitRepo(path string, f os.FileInfo, err error) error {
	if f.IsDir() && filepath.Base(path) == ".git" {
		dirs = append(dirs, path[:strings.LastIndex(path, "/")])
		return filepath.SkipDir
	}
	return nil
}

func gitCmd(idx int, path string, gitbinary string, action Action, p *parser.OutputParser) error {
	var out bytes.Buffer

	args := []string{gitbinary, action.Cmd}
	args = append(args, action.Args...)

	cmd := exec.Cmd{
		Dir:    path,
		Path:   gitbinary,
		Args:   args,
		Stdout: &out,
		Stderr: &out,
	}

	//cmd.Stdin = strings.NewReader("some input")

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR [%s][%s]\n", err, out)
		return err
	}

	dscr, err := p.Process(out, parser.NewDescriptor())
	//fmt.Fprintf(os.Stderr, "!!! [%s]\n", p.DebugOutput.String())

	if err != nil {
		fmt.Fprintf(os.Stderr, "\n<ERROR parser process>\n", path)
		fmt.Fprintf(os.Stderr, "=== %s ===\n", path)
		fmt.Fprintf(os.Stderr, "OUT:\n%s\n", out.String())
		fmt.Fprintf(os.Stderr, "ERR:\n%s\n", err.Error())
		fmt.Fprintf(os.Stderr, "DBG:\n%s\n", p.DebugOutput.String())
		fmt.Fprintf(os.Stderr, "</ERROR>\n")
	} else if p.Failed {
		fmt.Fprintf(os.Stderr, "\n<ERROR parser failed>\n", path)
		fmt.Fprintf(os.Stderr, "=== %s ===\n", path)
		fmt.Fprintf(os.Stderr, "OUT:\n%s\n", out.String())
		fmt.Fprintf(os.Stderr, "DBG:\n%s\n", p.DebugOutput.String())
		fmt.Fprintf(os.Stderr, "</ERROR>\n")
	} else {
		fmt.Fprintf(os.Stderr, "= %02d = %s %s", idx, path, dscr.AsString(dscr, true))
		// fmt.Printf("%s", out.String())
		// fmt.Printf(status.DebugOutput.String())
	}
	//fmt.Println("\n")
	return nil
}

func actionApplier(dirs []string, repos map[int]bool, action Action) error {
	var p *parser.OutputParser

	switch action.Cmd {
	case GIT_PULL.Cmd:
		p = pull.NewParser()
	case GIT_STATUS.Cmd:
		p = status.NewParser()
	case GIT_REBASE.Cmd:
		p = rebase.NewParser()
	case GIT_DIFF.Cmd:
		p = diff.NewParser()
	case GIT_CHECKOUT.Cmd:
		p = checkout.NewParser()
	case GIT_PUSH.Cmd:
		p = push.NewParser()
	case GIT_BRANCH.Cmd:
		p = branch.NewParser()
	default:
		return errors.New(fmt.Sprintf("unknown [%s] mode", action))
	}

	for i, d := range dirs {
		idx := i + 1
		l := len(repos)
		process := repos[idx]
		if l == 0 || process {
			if err := gitCmd(idx, d, gitbinary, action, p); err != nil {
				fmt.Fprintf(os.Stderr, "[%s] processing failed\n", d)
			}
		}
	}
	return nil
}

func walk(root string, repos map[int]bool, action Action) error {
	if err := filepath.Walk(root, walkerGetGitRepo); err != nil {
		return err
	}
	if err := actionApplier(dirs, repos, action); err != nil {
		return err
	}
	return nil
}

func init() {
	gb, err := exec.LookPath("git")
	checkError(err)
	gitbinary = gb
}

func main() {
	flags := map[string]interface{}{
		"root":  ".",
		"debug": false,
	}
	command := ""
	args := []string{}

	repos := map[int]bool{}

	for _, v := range os.Args[1:] {
		if command != "" {
			args = append(args, v)
		} else {
			if v[0] == '-' {
				parts := strings.SplitN(v, "=", -1)
				if len(parts) > 2 {
					fmt.Fprintf(os.Stderr, "flags parsing error [%s]\n", parts)
					return
				}
				switch parts[0] {
				default:
					fmt.Fprintf(os.Stderr, "unknown flag [%s]\n", parts)
				case "-r":
					for _, r := range strings.SplitN(parts[1], ",", -1) {
						val, _ := strconv.Atoi(r)
						repos[val] = true
					}
				case "--root":
					flags["root"] = parts[1]
				case "--debug":
					flags["debug"] = true
				}
			} else {
				command = v
			}
		}
	}

	switch {
	case flags["debug"]:
		fmt.Fprintf(os.Stderr, "debug mode")
		fmt.Fprintf(os.Stderr, "flags [%s]\n", flags)
		fmt.Fprintf(os.Stderr, "command [%s]\n", command)
		fmt.Fprintf(os.Stderr, "args [%s]\n", args)
	case command != "":
		action := Action{Cmd: command, Args: args}
		// fmt.Fprintf(os.Stderr, "started in [git %s] mode\n", action)
		if err := walk(flags["root"].(string), repos, action); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		}
	default:
		fmt.Fprintf(os.Stderr, "ERROR: unknown mode")
	}
}
