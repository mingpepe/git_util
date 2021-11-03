package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/mingpepe/git_util/util"
)

type CommitLog struct {
	Time     time.Time
	Msg      string
	FromPrev time.Duration
}

func parse(msg string) []CommitLog {
	layout := "2006-01-02 15:04:05"
	lines := strings.Split(msg, "\n")
	commits := make([]CommitLog, len(lines))
	for _, line := range lines {
		sep := strings.Split(line, "_")
		t, err := time.Parse(layout, sep[0][1:len(sep[0])-1])
		if err != nil {
			log.Panicf("Unexpected log format : %s, %v\n", sep[0][1:len(sep[0])-1], err)
		}
		c := CommitLog{
			Time: t,
			Msg:  sep[1],
		}
		commits = append(commits, c)
	}
	for i := 0; i < len(commits)-1; i++ {
		commits[i].FromPrev = commits[i].Time.Sub(commits[i+1].Time)
	}
	return commits
}

func (c *CommitLog) String() string {
	return fmt.Sprintf("%v %s takes %v", c.Time, c.Msg, c.FromPrev)
}

func main() {
	if !util.IsGitSupport() {
		log.Print("Git seems not installed yet")
		return
	}

	var path = flag.String("p", ".", "For git repo path")
	var n = flag.Int("n", -1, "Number of log")
	flag.Parse()

	if !util.IsGitDir(*path) {
		log.Print("Not a git repo")
		return
	}

	args := []string{
		"-C",
		*path,
		"log",
		"--pretty=format:%cd_%s",
		"--date=format:\"%Y-%m-%d %H:%M:%S\"",
	}
	if *n >= 1 {
		args = append(args, fmt.Sprintf("-n%d", *n))
	}

	cmd := exec.Command("git", args...)
	raw, err := cmd.Output()
	if err != nil {
		log.Panicf("Error while execute git log : %v\n", err)
	}
	for _, commit := range parse(string(raw)) {
		fmt.Println(commit.String())
	}
}
