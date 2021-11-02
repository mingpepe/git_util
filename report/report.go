package report

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

func IsGitSupport() bool {
	cmd := exec.Command("git", "version")
	raw, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.HasPrefix(string(raw), "git version")
}

func IsGitDir(path string) bool {
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

func (a *GitRepo) Parse(desc string) {
	a.Desc = desc
	a.BranchName = parseBranchName(desc)
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
	if IsGitDir(path) {
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
	cmd := exec.Command("git", "-C", path, "status")
	raw, err := cmd.Output()
	if err != nil {
		log.Panicf("Error while execute git status : %v\n", err)
	}
	a := GitRepo{Path: path}
	a.Parse(string(raw))
	return a
}

func parseBranchName(desc string) string {
	if strings.HasPrefix(desc, "On branch ") {
		firstLine := strings.Split(desc, "\n")[0]
		return strings.Split(firstLine, " ")[2]
	} else {
		log.Panicf("Unexpeced git log : %s\n", desc)
		return ""
	}
}
