package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	lines, err := readInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "read input: %v\n", err)
		os.Exit(1)
	}

	total := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		total += maxBankJoltage(line)
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

// maxBankJoltage picks two digits in order to form the largest possible two-digit number.
func maxBankJoltage(line string) int {
	digits := []rune(line)
	rightMax := make([]int, len(digits))

	// rightMax[i] holds the largest digit strictly to the right of i (or -1 if none).
	best := -1
	for i := len(digits) - 1; i >= 0; i-- {
		rightMax[i] = best
		if d := int(digits[i] - '0'); d > best {
			best = d
		}
	}

	maxVal := 0
	for i, tensRune := range digits {
		ones := rightMax[i]
		if ones == -1 {
			continue
		}
		tens := int(tensRune - '0')
		val := tens*10 + ones
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal
}
