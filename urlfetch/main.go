//go:build !solution

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Error")
		os.Exit(1)
	}

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		fmt.Println("Content of", url)
		fmt.Println("----------------------------------------")
		buf := make([]byte, 1024)
		for {
			n, err := resp.Body.Read(buf)
			if n > 0 {
				fmt.Print(string(buf[:n]))
			}
			if err != nil {
				break
			}
		}
		fmt.Println("----------------------------------------")
	}
}
