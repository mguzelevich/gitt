package actions

import (
	"errors"
)

type Cmd string

const (
	GIT_STATUS Cmd = "status"
)

type Action interface {
	Handle(path string)
	Output()
}

var actions = map[Cmd]Action{}

func init() {
	actions[GIT_STATUS] = GitStatus{}
}

func Check(cmd string) bool {
	return actions[Cmd(cmd)] != nil
}

func Get(cmd string) (Action, error) {
	if action, ok := actions[Cmd(cmd)]; ok {
		return action, nil
	}
	return nil, errors.New("err")
}
