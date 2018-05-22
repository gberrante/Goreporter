package main

import "os/exec"

func logger(path string) string {
	command := "--git-dir=" + path
	out, _ := exec.Command("git", command, "log", "--pretty=%aN;%aI").Output()
	return string(out[:])
}
