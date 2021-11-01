package main

import (
	"fmt"

	"github.com/mingpepe/git_util/util"
)

func main() {
	ret := util.Probe("C:\\Users\\user\\Desktop")
	for _, a := range ret {
		fmt.Println(a.String())
	}
}
