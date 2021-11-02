package util

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type GIT_STATE int

const (
	UPDATE_TO_DATE GIT_STATE = iota
	UN_PUSHED
	UN_COMMITED
	UN_STAGED
	NO_COMMITS_YET
)

func (state GIT_STATE) String() string {
	return [...]string{
		"UPDATE_TO_DATE",
		"UN_PUSHED",
		"UN_COMMITED",
		"UN_STAGED",
		"NO_COMMITS_YET",
	}[state]
}

func isGitDir(path string) bool {
	_, err := os.Stat(path + "\\.git")
	return !os.IsNotExist(err)
}

type GitRepo struct {
	Path              string
	Desc              string
	BranchName        string
	AnyUntrackedFiles bool
	State             GIT_STATE
}

func (a *GitRepo) Parse() {
	desc := a.Desc
	a.BranchName = strings.Split(strings.Split(desc, "\n")[0], " ")[2]
	a.AnyUntrackedFiles = strings.Contains(desc, "Untracked files:")
	if strings.Contains(desc, "up to date with") {
		a.State = UPDATE_TO_DATE
	} else if strings.Contains(desc, "Your branch is ahead of") {
		a.State = UN_PUSHED
	} else if strings.Contains(desc, "Changes to be committed:") {
		a.State = UN_COMMITED
	} else if strings.Contains(desc, "Changes not staged for commit") {
		a.State = UN_STAGED
	} else if strings.Contains(desc, "No commits yet") {
		a.State = NO_COMMITS_YET
	} else {
		log.Panicf("Unknown log msg : %s\n", desc)
	}
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
