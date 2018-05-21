package main

import "os"
import "fmt"

func main() {

	//argomento 0 è il path di questo exe, argomento 1 è quello inserito
	// i successivi argomenti separati dallo spazio sono ignorati
	path := os.Args[1]
	fmt.Println(path)
}
