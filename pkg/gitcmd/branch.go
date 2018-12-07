package gitcmd

import (
	"gopkg.in/src-d/go-git.v4"
)

// GetBranchName returns the git branch of the dir and error if there is an error.
func GetBranchName(dir string) (string, error) {
	r, err := git.PlainOpen(dir)
	if err != nil {
		return "", err
	}

	ref, err := r.Head()
	if err != nil {
		return "", err
	}

	refname := ref.Name()
	if err != nil {
		return "", err
	}

	return refname.Short(), nil
}
