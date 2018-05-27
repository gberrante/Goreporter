package main

import (
	"fmt"
)

// Generator of hash committer map, and other relative lists
func Generator(data []Commits) {
	NameMap := make(map[string]bool)

	for i := range data {
		if NameMap[data[i].name] == false {
			NameMap[data[i].name] = true
		}
	}
	fmt.Println("debug")
}
