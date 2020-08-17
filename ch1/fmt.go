package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	err := errors.New("unknown error")
	fmt.Printf("%+v\n", err)
	fmt.Println("hello")
}
