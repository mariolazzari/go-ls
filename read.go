package main

import (
	"fmt"
	"os"
	"strings"
)

func readPath(path string, showAll bool) []string {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// add . and .. to results
	var names []string
	if showAll {
		names = append(names, ".", "..")
	}

	for _, entry := range entries {
		// skip hidden files
		if !showAll && strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		names = append(names, entry.Name())
	}

	return names
}
