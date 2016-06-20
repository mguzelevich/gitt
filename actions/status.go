package actions

import (
	"fmt"
	"github.com/libgit2/git2go"
)

type GitStatus struct {
	a string
}

func (c GitStatus) Handle(path string) {
	repo, err := git.OpenRepository(path)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	fmt.Printf("%v\n", repo)

	opts := &git.StatusOptions{}
	opts.Show = git.StatusShowIndexAndWorkdir
	opts.Flags = git.StatusOptIncludeUntracked | git.StatusOptRenamesHeadToIndex | git.StatusOptSortCaseSensitively

	statusList, err := repo.StatusList(opts)
	if err != nil {
		fmt.Printf("%v", err)
	}

	entryCount, err := statusList.EntryCount()
	if err != nil {
		fmt.Printf("%v", err)
	}

	if entryCount != 1 {
		fmt.Printf("Incorrect number of status entries: %s\n", entryCount)
	}

	entry, err := statusList.ByIndex(0)
	if err != nil {
		fmt.Printf("%v", err)
	}

	if entry.Status != git.StatusWtNew {
		fmt.Printf("Incorrect status flags: %s\n", entry.Status)
	}
	if entry.IndexToWorkdir.NewFile.Path != "hello.txt" {
		fmt.Printf("Incorrect entry path: %s\n", entry.IndexToWorkdir.NewFile.Path)
	}
}

func (c GitStatus) Output() {

}
