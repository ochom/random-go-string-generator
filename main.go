package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		panic("usage: main.go <regex>")
	}

	pattern, err := NewPattern(args[0])
	if err != nil {
		panic("invalid pattern")
	}

	res := pattern.Generate()
	fmt.Println(res)
}
