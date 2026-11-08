package main

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"strconv"
	"syscall"
)

func printLong(dir, name string) {
	path := path.Join(dir, name)
	fileInfo, err := os.Lstat(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", name, err)
	}

	fileStats := fileInfo.Sys().(*syscall.Stat_t)

	owner, group := "?", "?"
	if u, error := user.LookupId(strconv.Itoa(int(fileStats.Uid))); error == nil {
		owner = u.Username
	}
	if g, error := user.LookupGroupId(strconv.Itoa(int(fileStats.Gid))); error == nil {
		group = g.Name
	}

	fmt.Printf("%s %d %s %s %d %s %s\n",
		modeString(fileStats.Mode),
		fileStats.Nlink,
		owner,
		group,
		fileInfo.Size(),
		fileInfo.ModTime().Format("Jan _2 15:04"),
		name,
	)
}

func printResults(names []string, path string, isLong bool) {
	for _, name := range names {
		if isLong {
			printLong(path, name)
		} else {
			fmt.Println(name)
		}
	}
}
