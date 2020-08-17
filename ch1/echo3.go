package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// a simple and more efficient solution would be to use the Join function from the strings package.
func main() {
	capacity := 100000
	var arr []string
	for i := 0; i < capacity; i++ {
		arr = append(arr, strconv.Itoa(i))
	}
	js := time.Now().Unix()
	strings.Join(arr, " ")
	je := time.Now().Unix()
	fmt.Printf("join function cost: %d\n", je - js)


	as := time.Now().Unix()
	sep := " "
	var s string
	for i := 0; i < capacity; i++ {
		s += sep + arr[i]
	}
	ae := time.Now().Unix()
	fmt.Printf("append function cost: %d\n", ae - as)
}
