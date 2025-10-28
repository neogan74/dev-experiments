package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	ignoreCase := flag.Bool("i", false, "ignore case")
	showLineNo := flag.Bool("n", false, "print line number")
	filePath := flag.String("f", "", "input file path(default: stdin)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage: %s [flags] <search>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "error: missing search argument")
		flag.Usage()
		os.Exit(2)
	}
	search := flag.Arg(0)
	if *ignoreCase {
		search = strings.ToLower(search)
	}

	var r io.Reader = os.Stdin
	if *filePath != "" {
		f, err := os.Open(*filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(2)
		}
		defer f.Close()
		r = f
	}

	sc := bufio.NewScanner(r)
	lineNo := 0
	matches := 0
	for sc.Scan() {
		line := sc.Text()
		target := line
		if *ignoreCase {
			target = strings.ToLower(line)
		}
		lineNo++
		if strings.Contains(target, search) {
			matches++
			if *showLineNo {
				fmt.Printf("%d: %s\n", lineNo, line)
			} else {
				fmt.Println(line)
			}
		}
	}
	// foo
	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	if matches == 0 {
		os.Exit(1)
	}
}
