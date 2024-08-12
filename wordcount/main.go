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
		fmt.Println("Usage two files")
		os.Exit(1)
	}

	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wordcount: %v\n", err)
			continue
		}
		defer file.Close()
		countsLines(file, counts)
	}
	printCount(counts)
}

func countsLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[string(input.Bytes())]++
	}
}

type kv struct {
	Key   string
	Value int
}

func printCount(count map[string]int) {
	var sortedCounts []kv
	for k, v := range count {
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
