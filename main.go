package main

import (
	"fmt"
	"sort"
)

func printSorted(args map[int]string) {
	keys := make([]int, 0, len(args))
	for k := range args {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var sorted []string
	for _, i := range keys {
		sorted = append(sorted, args[i])
	}
	fmt.Println(sorted)
}

func main() {
	printSorted(map[int]string{2:"a", 0:"b", 1:"c"})
	printSorted(map[int]string{10: "aa", 0: "bb", 500: "cc"})
}