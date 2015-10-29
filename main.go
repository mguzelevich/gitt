package main

import (
	"bytes"
	//	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mguzelevich/gitt/git/pull"
	"github.com/mguzelevich/gitt/git/status"
	"github.com/mguzelevich/gitt/parser"
)

type Action string

const (
	GIT_PULL   Action = "pull"
	GIT_STATUS Action = "status"
)

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

func gitCmd(path string, gitbinary string, cmdString string, p *parser.OutputParser) error {
	cmd := exec.Command(gitbinary, cmdString)
	cmd.Dir = path

	//cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out // os.Stdout
	cmd.Stderr = &out // os.Stderr

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
		fmt.Fprintf(os.Stderr, "=== %s %s", path, dscr.AsString(dscr, true))
		// fmt.Printf("%s", out.String())
		// fmt.Printf(status.DebugOutput.String())
	}
	//fmt.Println("\n")
	return nil
}

func walker(root string, action Action) {
	var f filepath.WalkFunc

	switch action {
	case GIT_PULL:
		f = walkerGetGitRepo
	case GIT_STATUS:
		f = walkerGetGitRepo
	default:
		fmt.Fprintf(os.Stderr, "unknown [%s] mode", action)
	}

	err := filepath.Walk(root, f)
	checkError(err)
}

func actionApplier(dirs []string, action Action) {
	var cmd string
	var p *parser.OutputParser

	switch action {
	case GIT_PULL:
		cmd = "pull"
		p = pull.NewParser()
	case GIT_STATUS:
		cmd = "status"
		p = status.NewParser()
	default:
		fmt.Fprintf(os.Stderr, "unknown [%s] mode", action)
	}

	for _, d := range dirs {
		if err := gitCmd(d, gitbinary, cmd, p); err != nil {
			fmt.Fprintf(os.Stderr, "[%s] processing failed\n", d)
		}
	}
}

func walk(root string, action Action) {
	walker(root, action)
	actionApplier(dirs, action)
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
		action := Action(command)
		fmt.Fprintf(os.Stderr, "started in [git %s] mode\n\n", action)
		walk(flags["root"].(string), action)
	default:
		fmt.Fprintf(os.Stderr, "unknown mode")
	}
}
