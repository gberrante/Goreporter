package main

import (
	"fmt"
)

type Counter struct {
	exist  bool
	numCom int
}

// Generator of committer map, and other relative lists
func Generator(data []Commits) {
	//Committer HASH Table
	NameMap := make(map[string]Counter)
	DateMap := make(map[string]Counter)
	WeekMap := make(map[string]Counter)

	for i := range data {
		if NameMap[data[i].name].exist == false {
			value := NameMap[data[i].name].numCom
			NameMap[data[i].name] = Counter{exist: true}
			NameMap[data[i].name] = Counter{numCom: value + 1}
		}
		if DateMap[data[i].date].exist == false {
			value := DateMap[data[i].date].numCom
			DateMap[data[i].date] = Counter{exist: true}
			DateMap[data[i].date] = Counter{numCom: value + 1}
		}
		if WeekMap[data[i].week].exist == false {
			value := WeekMap[data[i].week].numCom
			WeekMap[data[i].week] = Counter{exist: true}
			WeekMap[data[i].week] = Counter{numCom: value + 1}
		}
	}

	fmt.Println("debug")
}
