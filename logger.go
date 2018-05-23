package main

import (
	"os/exec"
	"regexp"
	"strconv"
)

//commit data structure
type Commits struct {
	name  string
	year  int
	month int
	day   int
}

func logger(path string) []Commits {
	command := "--git-dir=" + path
	out, _ := exec.Command("git", command, "log", "--pretty=%an;%ai").Output()
	text := string(out[:])
	var expr = regexp.MustCompile(`(.*)\n`)
	lines := expr.FindStringSubmatch(text)
	imax := len(lines)
	list := []Commits{}

	for i := 0; i < imax; i++ {
		token := regexp.MustCompile(`^(.+);(\d+)-(\d+)-(\d+)`).FindStringSubmatch(lines[i])
		temp := Commits{}
		temp.name = token[1]
		temp.year, _ = strconv.Atoi(token[2])
		temp.month, _ = strconv.Atoi(token[3])
		temp.day, _ = strconv.Atoi(token[4])
		list = append(list, temp)
	}
	return list
}
