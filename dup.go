// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    counts := make(map[string]int)
	appearances := make(map[string][]string)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts, appearances)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts, appearances)
            f.Close()
        }
    }
    for line, n := range counts {
        if n > 1 {
			fileAppearances := ""
			for _, fileName := range appearances[line] {
				fileAppearances = fileAppearances + " " + fileName
			}
            fmt.Printf("%d\t%s\t%s\n", n, line, fileAppearances)
        }
    }
}

func countLines(f *os.File, counts map[string]int, appearances map[string][]string) {
    input := bufio.NewScanner(f)

    for input.Scan() {
        counts[input.Text()]++
		appearances[input.Text()] = addAppearance(appearances[input.Text()], f.Name())
    }
    // NOTE: ignoring potential errors from input.Err()
}

func addAppearance(appearances []string, fileName string) []string{
	for _, value := range appearances {
		if value == fileName {
			return appearances
		}
	}
	appearances = append(appearances, fileName)
	return appearances
}
