package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mguzelevich/gitt/parser"
	"github.com/mguzelevich/gitt/pull"
	"github.com/mguzelevich/gitt/status"
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

func gitCmd(path string, cmd *exec.Cmd, p *parser.OutputParser) {
	//	cmd := exec.Command(gitbinary, "pull")
	cmd.Dir = path

	//cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out // os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Err %q\n", err)
		return
	}

	dscr := parser.NewDescriptor()

	if dscr, err := p.Process(out, dscr); err != nil {
		fmt.Fprintf(os.Stderr, "=== %s ===\n", path)
		fmt.Printf("%s", out.String())
		fmt.Printf("Err %s", err.Error())
		fmt.Printf(p.DebugOutput.String())
	} else if p.Failed {
		fmt.Fprintf(os.Stderr, "=== %s ===\n", path)
		fmt.Printf("%s", out.String())
		fmt.Printf("Err %s", err.Error())
		fmt.Printf(p.DebugOutput.String())
	} else {
		fmt.Fprintf(os.Stderr, "=== %s %s", path, dscr.AsString(dscr, true))
		// fmt.Printf("%s", out.String())
		// fmt.Printf(status.DebugOutput.String())
	}
	//fmt.Println("\n")
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
		gitCmd(d, exec.Command(gitbinary, cmd), p)
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
	//debug := flag.Bool("debug", false, "debug mode")
	git := flag.String("git", "status", "git action")
	flag.Parse()

	root := flag.Arg(0)
	if root == "" {
		root = "."
	}

	switch {
	// case *debug:
	// 	fmt.Fprintf(os.Stderr, "debug mode")
	case *git != "":
		action := Action(*git)
		fmt.Fprintf(os.Stderr, "started in [git %s] mode\n\n", action)
		walk(root, action)
	default:
		fmt.Fprintf(os.Stderr, "unknown mode")
	}
}
