package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/mingpepe/git_util/util"
)

func printGitRepo(r util.GitRepo) {
	msg := fmt.Sprintf("%-50s %-20s %-20s", r.Path, r.BranchName, r.State.String())
	switch r.State {
	case util.UPDATE_TO_DATE:
		color.Green(msg)
	case util.UN_PUSHED:
		color.Yellow(msg)
	case util.UN_COMMITED:
		color.Red(msg)
	case util.UN_STAGED:
		color.Red(msg)
	case util.NO_COMMITS_YET:
		color.Red(msg)
	default:
		panic("unknown git state")
	}
}

func main() {
	var path = flag.String("p", ".", "For git repo path")
	flag.Parse()

	if util.IsGitSupport() {
		ret := util.Probe(*path)
		for _, r := range ret {
			printGitRepo(r)
		}
	} else {
		log.Print("Git seems not installed yet")
	}

}
