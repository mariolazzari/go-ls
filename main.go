package main

func main() {
	path, showAll, isLong := parser()
	entries := readPath(path, showAll)
	printResults(entries, path, isLong)
}
