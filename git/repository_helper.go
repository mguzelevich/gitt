package git

import (
//	"errors"
)

type RepositoryHelper struct {
	repository Repository
	status     bool
}

func (r *RepositoryHelper) __init__(RepositoryHelper string) {

}

func (r *RepositoryHelper) Status() bool {
	return r.status
}

func (r *RepositoryHelper) checkout(branch string) {

}

func (r *RepositoryHelper) clone(target string) {

}

func (r *RepositoryHelper) merge(branch string) {

}

func (r *RepositoryHelper) push(args string) {

}

func (r *RepositoryHelper) tag(tag string) {

}
