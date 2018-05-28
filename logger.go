package main

import (
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

//Commits data structure
type Commits struct {
	name     string
	date     string
	week     string
	add      int
	delete   int
	filesMod int
}

func reader(path string) []Commits {
	command := "--git-dir=" + path
	out, _ := exec.Command("git", command, "log", "--pretty=#BEGIN%an;%ai;%aD", "--numstat").Output()
	text := regexp.MustCompile(`\n\n`).ReplaceAllString(string(out[:]), "\n")
	blocks := strings.Split(text, "#BEGIN")
	var list []Commits

	for i := 1; i < len(blocks); i++ {
		lines := strings.Split(blocks[i], "\n")
		lines = lines[:len(lines)-1]

		token := regexp.MustCompile(`^(.+);(\S+)\s.*;(\w+),`).FindStringSubmatch(lines[0])
		temp := Commits{}
		temp.name = token[1]
		temp.date = token[2]
		temp.week = token[3]

		added := 0
		removed := 0
		modified := 0

		for j := 1; j < len(lines); j++ {
			expr, _ := regexp.Compile(`(\d+)\s(\d+)\s.*`)
			if expr.MatchString(lines[j]) {
				token := expr.FindStringSubmatch(lines[j])
				adj, _ := strconv.Atoi(token[1])
				rmj, _ := strconv.Atoi(token[2])
				added += adj
				removed += rmj
				modified++
			} else {
				continue
			}
		}
		temp.filesMod = modified
		temp.add = added
		temp.delete = removed
		list = append(list, temp)
	}
	return list
}
