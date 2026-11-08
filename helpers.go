package main

import (
	"os"
	"syscall"
)

func modeString(mode uint16) string {
	var s [10]byte
	switch mode & syscall.S_IFMT {
	case syscall.S_IFDIR:
		s[0] = 'd'
	case syscall.S_IFLNK:
		s[0] = 'l'
	case syscall.S_IFCHR:
		s[0] = 'c'
	case syscall.S_IFBLK:
		s[0] = 'b'
	case syscall.S_IFIFO:
		s[0] = 'p'
	case syscall.S_IFSOCK:
		s[0] = 's'
	default:
		s[0] = '-'
	}
	return string(s[:1]) + os.FileMode(mode).Perm().String()[1:]

}
