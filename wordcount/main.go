//go:build !solution

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wordcount file1 file2 ...")
		os.Exit(1)
	}

	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wordcount: %v\n", err)
			continue
		}
		countLines(file, counts)
		file.Close()
	}

	printCounts(counts)
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

func printCounts(counts map[string]int) {
	type kv struct {
		Key   string
		Value int
	}

	var sortedCounts []kv
	for k, v := range counts {
		if v >= 2 {
			sortedCounts = append(sortedCounts, kv{k, v})
		}
	}

	sort.Slice(sortedCounts, func(i, j int) bool {
		return sortedCounts[i].Value > sortedCounts[j].Value
	})

	for _, kv := range sortedCounts {
		fmt.Printf("%d\t%s\n", kv.Value, kv.Key)
	}
}
