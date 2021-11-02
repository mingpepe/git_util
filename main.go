package main

import (
	"flag"
	"log"

	"github.com/mingpepe/git_util/util"
)

func main() {
	var path = flag.String("p", ".", "For git repo path")
	flag.Parse()

	if util.IsGitSupport() {
		ret := util.Probe(*path)
		for _, a := range ret {
			log.Printf("Path : %s\n", a.Path)
			log.Printf("Branch name : %s\n", a.BranchName)
			log.Printf("Any untracked files : %v\n", a.AnyUntrackedFiles)
			log.Printf("State : %v\n\n", a.State)
		}
	} else {
		log.Print("Git seems not installed yet")
	}

}
