package main

import (
	"fmt"
	"os"
	"strings"
)

func parser() (string, bool, bool) {

	path, showAll, isLong := ".", false, false

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
			case 'l':
				isLong = true
			default:
				fmt.Fprintf(os.Stderr, "usage: %s [-al] [path]\n", os.Args[0])
				os.Exit(1)
			}
		}
	}

	return path, showAll, isLong
}
