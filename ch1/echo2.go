package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for index, arg := range os.Args[1:] {
		fmt.Printf("index: %d, value: %s\n", index, arg)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
