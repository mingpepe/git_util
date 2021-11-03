package util

import (
	"os"
	"os/exec"
	"strings"
)

func IsGitSupport() bool {
	cmd := exec.Command("git", "version")
	raw, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.HasPrefix(string(raw), "git version")
}

func IsGitDir(path string) bool {
	_, err := os.Stat(path + "//.git")
	return !os.IsNotExist(err)
}
