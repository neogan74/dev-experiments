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
	partTwo := flag.Bool("part2", false, "compute part two (use 12 batteries per bank)")
	flag.Parse()

	pick := 2
	if *partTwo {
		pick = 12
	}

	lines, err := readInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "read input: %v\n", err)
		os.Exit(1)
	}

	var total int64
	for _, line := range lines {
		if line == "" {
			continue
		}
		total += maxBankJoltage(line, pick)
	}

	fmt.Println(total)
}

// readInput reads from stdin if data is present; otherwise falls back to input.txt.
func readInput() ([]string, error) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return readLines(os.Stdin)
	}

	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return readLines(f)
}

func readLines(r io.Reader) ([]string, error) {
	var lines []string
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		lines = append(lines, strings.TrimSpace(sc.Text()))
	}
	return lines, sc.Err()
}

// maxBankJoltage picks `pick` digits in order to form the largest possible number.
func maxBankJoltage(line string, pick int) int64 {
	if pick <= 0 {
		return 0
	}

	digits := []byte(line)
	if pick > len(digits) {
		pick = len(digits)
	}

	// Greedy monotonic stack: drop smaller leading digits when there are enough remaining.
	stack := make([]byte, 0, len(digits))
	remaining := len(digits)

	for _, d := range digits {
		remaining--
		for len(stack) > 0 && stack[len(stack)-1] < d && len(stack)+remaining >= pick {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, d)
	}

	if len(stack) > pick {
		stack = stack[:pick]
	}

	var val int64
	for _, d := range stack {
		val = val*10 + int64(d-'0')
	}

	return val
}
