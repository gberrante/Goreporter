package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var Authors [][]string
var Days [][]string
var Weeks [][]string

// Generator of committer map, and other relative lists
func Generator(data []Commits) {
	//Committer HASH Table
	NameMap := make(map[string]int)
	DateMap := make(map[string]int)
	WeekMap := make(map[string]int)

	//Scan Data and remove duplicates
	for i := range data {
		NameMap[data[i].name]++
		DateMap[data[i].date]++
		WeekMap[data[i].week]++
	}
	jsonName, _ := json.Marshal(NameMap)
	jsonDate, _ := json.Marshal(DateMap)
	jsonWeek, _ := json.Marshal(WeekMap)
	file, err := os.Create("result.json")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	fmt.Fprintf(file, "%s\n%s\n%s", jsonName, jsonDate, jsonWeek)
}
