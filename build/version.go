package build

import (
	git "gopkg.in/libgit2/git2go.v24"
)

func GetVersion() (string, error) {
	repo, err := git.OpenRepository(".")
	if err != nil {
		return "", err
	}
	defer repo.Free()
	head, err := repo.Head()
	if err != nil {
		return "", err
	}
	defer head.Free()
	oid := head.Target()
	commit, err := repo.LookupCommit(oid)
	if err != nil {
		return "", err
	}
	opts, err := git.DefaultDescribeOptions()
	if err != nil {
		return "", err
	}
	opts.Strategy = git.DescribeTags
	opts.Pattern = "v[0-9.]*[a-z]"
	opts.ShowCommitOidAsFallback = true
	desc, err := commit.Describe(&opts)
	if err != nil {
		return "", err
	}
	defer desc.Free()
	fmtOpts, err := git.DefaultDescribeFormatOptions()
	if err != nil {
		return "", err
	}
	return desc.Format(&fmtOpts)
}
