package main

import (
	"flag"
	"fmt"

	"github.com/mingpepe/git_util/util"
)

func main() {
	var path = flag.String("p", ".", "For git repo path")
	flag.Parse()

	ret := util.Probe(*path)
	for _, a := range ret {
		fmt.Println(a.String())
	}
}
