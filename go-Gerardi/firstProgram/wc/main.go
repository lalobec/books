package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Define a flag, "-l", to count the lines instead of words
	// default value is false
	lines := flag.Bool("l", false, "Count lines")

	// Parse (analyze) the flags provided by user
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r)

	// If the count lines flag is false, we count the words:
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
