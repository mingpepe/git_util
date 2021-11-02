package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/mingpepe/git_util/report"
	"github.com/mingpepe/git_util/util"
)

const FMT = "%-50s %-20s %-20s"

func printGitRepoTitle() {
	msg := fmt.Sprintf(FMT, "Path", "BranchName", "State")
	color.Green(msg)
}

func printGitRepo(r report.GitRepo) {
	msg := fmt.Sprintf(FMT, r.Path, r.BranchName, r.State.String())
	switch r.State {
	case report.UPDATE_TO_DATE:
		color.Green(msg)
	case report.UN_PUSHED:
		color.Yellow(msg)
	case report.UN_COMMITED:
		color.Red(msg)
	case report.UN_STAGED:
		color.Red(msg)
	case report.NO_COMMITS_YET:
		color.Red(msg)
	default:
		panic("unknown git state")
	}
}

func main() {
	var path = flag.String("p", ".", "For git repo path")
	flag.Parse()

	if util.IsGitSupport() {
		ret := report.Probe(*path)
		printGitRepoTitle()
		for _, r := range ret {
			printGitRepo(r)
		}
	} else {
		log.Print("Git seems not installed yet")
	}
}
