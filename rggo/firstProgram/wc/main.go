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

	// Second flag, "-b", to count bytes
	byte_count := flag.Bool("b", false, "Count bytes")

	// Parse (analyze) the flags provided by user
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *byte_count))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines {
		if !countBytes {
			scanner.Split(bufio.ScanWords)
		} else {
			scanner.Split(bufio.ScanBytes)
		}
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
