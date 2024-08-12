//go:build !solution

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	var mu sync.Mutex
	results := make([]string, 0)

	for _, url := range os.Args[1:] {
		wg.Add(1)
		go fetch(url, &wg, &mu, &results)
	}

	wg.Wait()

	for _, result := range results {
		fmt.Println(result)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, wg *sync.WaitGroup, mu *sync.Mutex, results *[]string) {
	defer wg.Done()

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		mu.Lock()
		*results = append(*results, fmt.Sprintf("Get %s: %v", url, err))
		mu.Unlock()
		return
	}
	defer resp.Body.Close()

	nbytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		mu.Lock()
		*results = append(*results, fmt.Sprintf("while reading %s: %v", url, err))
		mu.Unlock()
		return
	}

	secs := time.Since(start).Seconds()
	mu.Lock()
	*results = append(*results, fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url))
	mu.Unlock()
}
