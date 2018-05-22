package main

import (
	"fmt"
	"os/exec"
	"regexp"
)

func logger(path string) string {
	command := "--git-dir=" + path
	out, _ := exec.Command("git", command, "log", "--pretty=%an;%ai").Output()
	text := string(out[:])
	var expr = regexp.MustCompile(`\n`)
	lines := expr.FindStringSubmatch(text)
	imax := len(lines)
	for i := 1; i < imax; i++ {
		fmt.Println(lines[i])
	}
	//var expr = regexp.MustCompile(`^(.*);(\d+)-(\d+)-(\d+)`)
	//token := expr.FindStringSubmatch(string(out[:]))
	return "fine"
}
