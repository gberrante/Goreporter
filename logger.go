package main

import (
	"fmt"
	"os/exec"
	"strings"
)

//commit data structure
type Commits struct {
	name   string
	year   int
	month  int
	day    int
	add    int
	delete int
}

func logger(path string) int {
	command := "--git-dir=" + path
	out, _ := exec.Command("git", command, "log", "--pretty=#BEGIN%an;%ai", "--numstat").Output()
	text := string(out[:])
	blocks := strings.SplitAfter(text, "#BEGIN")
	imax := len(blocks)
	blocks = blocks[1:]
	for i := 0; i < imax; i++ {
		lines := strings.SplitAfter(blocks[i], "\n")
		lines = lines[:len(lines)-1]
		lines = append(lines[:1], lines[2:]...)
		fmt.Println(lines[0])
		/*token := regexp.MustCompile(`^#BEGIN(.+);(\d+)-(\d+)-(\d+)`).FindStringSubmatch(lines[i])
		temp := Commits{}
		temp.name = token[1]
		temp.year, _ = strconv.Atoi(token[2])
		temp.month, _ = strconv.Atoi(token[3])
		temp.day, _ = strconv.Atoi(token[4])
		list = append(list, temp)*/
	}
	return 0
}
