package main

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

	//Scan Data and remove duplicates
	for i := range data {
		if !NameMap[data[i].name].exist {
			NameMap[data[i].name] = Counter{exist: true}
		}
		if !DateMap[data[i].date].exist {
			DateMap[data[i].date] = Counter{exist: true}
		}
		if !WeekMap[data[i].week].exist {
			WeekMap[data[i].week] = Counter{exist: true}
		}
	}
	//Count data occurence for each item in Map
	for i := range data {
		NameMap[data[i].name] = Counter{exist: true, numCom: NameMap[data[i].name].numCom + 1}
		DateMap[data[i].date] = Counter{exist: true, numCom: DateMap[data[i].date].numCom + 1}
		WeekMap[data[i].week] = Counter{exist: true, numCom: WeekMap[data[i].week].numCom + 1}
	}
	/*
		NameK := ExtractKey(NameMap)
		DateK := ExtractKey(DateMap)
		WeeK := ExtractKey(WeekMap)*/

}

func ExtractKey(table map[string]Counter) []string {
	var keys []string
	for key, _ := range table {
		keys = append(keys, key)
	}
	return keys
}
