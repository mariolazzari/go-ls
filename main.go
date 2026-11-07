package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	path := "."
	showAll := false

	// args parser
	for _, arg := range os.Args[1:] {
		// no - found
		if !strings.HasPrefix(arg, "-") {
			path = arg
			continue
		}

		// search for '-'
		for _, char := range arg[1:] {
			switch char {
			case 'a':
				showAll = true
			default:
				fmt.Fprintf(os.Stderr, "usage: %s [-a] [path]\n", os.Args[0])
				os.Exit(1)
			}
		}
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, entry := range entries {
		// skip hidden files
		if !showAll && strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		fmt.Println(entry.Name())
	}
}
