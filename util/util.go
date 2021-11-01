package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func isGitDir(path string) bool {
	_, err := os.Stat(path + "\\.git")
	return !os.IsNotExist(err)
}

type GitRepo struct {
	Path                 string
	Desc                 string
	AnyUntrackedFiles    bool
	ChangesToBeCommitted bool
	NotFullPush          bool
}

func (a *GitRepo) String() string {
	return fmt.Sprintf("%s\nAnyUntrackedFiles : %v\nChangesToBeCommitted:%v\nNotFullPush:%v\n", a.Path, a.AnyUntrackedFiles, a.ChangesToBeCommitted, a.NotFullPush)
}

func (a *GitRepo) Parse() {
	a.AnyUntrackedFiles = strings.Contains(a.Desc, "Untracked files:")
	a.ChangesToBeCommitted = strings.Contains(a.Desc, "Changes to be committed:")
	a.NotFullPush = !strings.Contains(a.Desc, "Your branch is ahead of")
}

func Probe(path string) []GitRepo {
	ret, _ := probeInternal(path)
	return ret
}

func probeInternal(path string) ([]GitRepo, bool) {
	ret := make([]GitRepo, 0)
	if isGitDir(path) {
		ret = append(ret, checkStatus(path))
		return ret, true
	} else {
		fileInfo, _ := ioutil.ReadDir(path)
		any := false
		for _, file := range fileInfo {
			if file.IsDir() {
				r, a := probeInternal(path + "\\" + file.Name())
				if a {
					ret = append(ret, r...)
					any = true
				}
			}
		}
		return ret, any
	}
}

func checkStatus(path string) GitRepo {
	a := GitRepo{Path: path}
	c, b := exec.Command("git", "-C", path, "status"), new(strings.Builder)
	c.Stdout = b
	c.Run()

	a.Desc = b.String()
	a.Parse()
	return a
}
