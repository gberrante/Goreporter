package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	//argomento 0 è il path di questo exe, argomento 1 è quello inserito
	// i successivi argomenti separati dallo spazio sono ignorati
	//path := os.Args[1]
	//fmt.Println(path)
	out, err := exec.Command("git", "--version").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)
}
