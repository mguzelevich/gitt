package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/libgit2/git2go"
	// "github.com/mguzelevich/gitt/git/branch"
	// "github.com/mguzelevich/gitt/git/checkout"
	// "github.com/mguzelevich/gitt/git/commit"
	// "github.com/mguzelevich/gitt/git/diff"
	// "github.com/mguzelevich/gitt/git/fetch"
	// "github.com/mguzelevich/gitt/git/pull"
	// "github.com/mguzelevich/gitt/git/push"
	// "github.com/mguzelevich/gitt/git/rebase"

	"github.com/mguzelevich/gitt/actions"
)

var repos Repos

type Repo struct {
	repo    *git.Repository
	path    string
	exclude bool
}

func (r *Repo) String() string {
	return fmt.Sprintf("%s %s", r.exclude, r.path)
}

type Repos struct {
	repos   []Repo
	process map[int]bool
}

func (repos *Repos) walker(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		if r, err := git.OpenRepository(path); err != nil {
			// fmt.Printf("path != repo: %s skipped %s\n", path, err)
		} else {
			// fmt.Printf("repo: %s (%s)\n", path, r)
			repos.repos = append(repos.repos, Repo{repo: r, path: path, exclude: false})
			return filepath.SkipDir
		}
	}
	return nil
}

func (repos *Repos) walk(root string) error {
	if err := filepath.Walk(root, repos.walker); err != nil {
		return err
	}
	for i, _ := range repos.repos {
		idx := i + 1
		process := repos.process[idx]
		repos.repos[i].exclude = !(len(repos.process) == 0 || process)
	}
	return nil
}

func (repos *Repos) output() error {
	for i, r := range repos.repos {
		idx := i + 1
		fmt.Printf("%05d: %v\n", idx, r)
	}
	return nil
}

func (repos *Repos) handle(action actions.Action) error {
	for i := range repos.repos {
		r := repos.repos[i]
		if !r.exclude {
			action.Handle(r.path)
		}
	}
	return nil
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}

// func walk(root string, repos map[int]bool) error {
// 	d := Dirs{}
// 	if err := filepath.Walk(root, d.walk); err != nil {
// 		return err
// 	}
// 	return nil
// }

func init() {
	repos = Repos{repos: []Repo{}, process: map[int]bool{}}
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
				case "-r":
					for _, r := range strings.SplitN(parts[1], ",", -1) {
						if rr := strings.SplitN(r, "-", -1); len(rr) > 1 {
							start, _ := strconv.Atoi(rr[0])
							finish, _ := strconv.Atoi(rr[1])
							if finish <= start {
								continue
							}
							for i := start; i <= finish; i++ {
								repos.process[i] = true
							}
						} else {
							val, _ := strconv.Atoi(r)
							repos.process[val] = true
						}
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

	if !actions.Check(command) {
		panic("unknown command")
	}

	switch {
	case flags["debug"]:
		fmt.Fprintf(os.Stderr, "debug mode")
		// fmt.Fprintf(os.Stderr, "flags [%s]\n", flags)
		// fmt.Fprintf(os.Stderr, "command [%s]\n", command)
		// fmt.Fprintf(os.Stderr, "args [%s]\n", args)
	case command != "":
		// 	action := Action{Cmd: command, Args: args}
		// 	// fmt.Fprintf(os.Stderr, "started in [git %s] mode\n", action)
	default:
		fmt.Fprintf(os.Stderr, "ERROR: unknown mode")
	}

	if err := repos.walk(flags["root"].(string)); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	} else {
		// repos.output()
		if action, err := actions.Get(command); err == nil {
			//panic("unknown")
			repos.handle(action)
		}

	}

	// repo, err := git.Clone("https://github.com/mguzelevich/gitt.git", "/tmp/gitt-tmp", &git.CloneOptions{})
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Fprintf(os.Stderr, "repo [%s]\n", repo)
	// }
}
