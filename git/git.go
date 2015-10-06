package git

import (
	"errors"
)

type BranchesLink struct {
	Hash   string
	Local  string
	Remote string
}

type FileData struct {
	Filename     string
	OldFilename  string
	DiffPercent  int
	Lines        int
	LinesAdded   int
	LinesDeleted int
	Operation    string
	Mode         int
}

type Files []FileData

func (slice Files) Len() int {
	return len(slice)
}

func (slice Files) Less(i, j int) bool {
	return slice[i].Filename < slice[j].Filename
}

func (slice Files) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func ParseDiffLine(line string) (int, int, error) {
	add := 0
	del := 0
	for _, s := range line {
		switch s {
		default:
			return 0, 0, errors.New("unknown diff line symbol")
		case '+':
			add += 1
		case '-':
			del += 1
		}
	}
	return add, del, nil
}
