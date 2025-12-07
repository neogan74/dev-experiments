package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

const defaultInput = "input.txt"

type interval struct {
	start int
	end   int
}

func parseRange(line string) (interval, error) {
	parts := strings.Split(strings.TrimSpace(line), "-")
	if len(parts) != 2 {
		return interval{}, fmt.Errorf("invalid range %q", line)
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return interval{}, fmt.Errorf("invalid range start %q: %w", parts[0], err)
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return interval{}, fmt.Errorf("invalid range end %q: %w", parts[1], err)
	}

	if start > end {
		start, end = end, start
	}

	return interval{start: start, end: end}, nil
}

func parseInt(line string) (int, error) {
	value, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		return 0, fmt.Errorf("invalid number %q: %w", line, err)
	}
	return value, nil
}

func isFresh(id int, ranges []interval) bool {
	for _, r := range ranges {
		if id >= r.start && id <= r.end {
			return true
		}
	}
	return false
}

func loadInput(r io.Reader) ([]interval, []int, error) {
	scanner := bufio.NewScanner(r)
	// Allow moderately large lines.
	const maxLine = 1024 * 1024
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, maxLine)

	var ranges []interval
	var ids []int

	readingRanges := true
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			readingRanges = false
			continue
		}

		if readingRanges {
			r, err := parseRange(line)
			if err != nil {
				return nil, nil, err
			}
			ranges = append(ranges, r)
			continue
		}

		id, err := parseInt(line)
		if err != nil {
			return nil, nil, err
		}
		ids = append(ids, id)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return ranges, ids, nil
}

func loadInputFile(path string) ([]interval, []int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	return loadInput(file)
}

func countFreshIDs(ids []int, ranges []interval) int {
	freshCount := 0
	for _, id := range ids {
		if isFresh(id, ranges) {
			freshCount++
		}
	}
	return freshCount
}

func countFreshRangeIDs(ranges []interval) int {
	if len(ranges) == 0 {
		return 0
	}

	sorted := make([]interval, len(ranges))
	copy(sorted, ranges)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].start < sorted[j].start
	})

	total := 0
	current := sorted[0]
	for _, r := range sorted[1:] {
		if r.start <= current.end+1 {
			if r.end > current.end {
				current.end = r.end
			}
			continue
		}

		total += current.end - current.start + 1
		current = r
	}

	total += current.end - current.start + 1
	return total
}

func main() {
	var part2 bool

	root := &cobra.Command{
		Use:   "day5 [input]",
		Short: "Count fresh ingredient IDs (part 1) or total fresh IDs across ranges (part 2)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := defaultInput
			if len(args) == 1 {
				path = args[0]
			}

			ranges, ids, err := loadInputFile(path)
			if err != nil {
				return err
			}

			if part2 {
				fmt.Println(countFreshRangeIDs(ranges))
				return nil
			}

			fmt.Println(countFreshIDs(ids, ranges))
			return nil
		},
	}

	root.Flags().BoolVar(&part2, "part2", false, "run part two calculation")

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
