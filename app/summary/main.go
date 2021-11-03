package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/mingpepe/git_util/repo"
	"github.com/mingpepe/git_util/util"
)

const FMT = "%-50s %-20s %-20s"

func printGitRepoTitle() {
	msg := fmt.Sprintf(FMT, "Path", "BranchName", "State")
	color.Green(msg)
}

func printGitRepo(r repo.GitRepo) {
	msg := fmt.Sprintf(FMT, r.Path, r.BranchName, r.State.String())
	switch r.State {
	case repo.UPDATE_TO_DATE:
		color.Green(msg)
	case repo.UN_PUSHED:
		color.Yellow(msg)
	case repo.UN_COMMITED:
		color.Red(msg)
	case repo.UN_STAGED:
		color.Red(msg)
	case repo.NO_COMMITS_YET:
		color.Red(msg)
	case repo.UNKNOWN:
		color.Magenta(msg)
	default:
		panic("unknown git state")
	}
}

func main() {
	var path = flag.String("p", ".", "For git repo path")
	flag.Parse()

	if util.IsGitSupport() {
		ret := repo.Probe(*path)
		printGitRepoTitle()
		for _, r := range ret {
			printGitRepo(r)
		}
	} else {
		log.Print("Git seems not installed yet")
	}
}
