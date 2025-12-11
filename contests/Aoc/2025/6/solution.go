package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2 := flag.Bool("part2", false, "enable part 2 interpretation (columns right-to-left)")
	flag.Parse()

	lines, width, err := readLines()
	if err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
		os.Exit(1)
	}

	if len(lines) == 0 {
		fmt.Println(0)
		return
	}

	grid := padLines(lines, width)
	total, err := solveGrid(grid, *part2)
	if err != nil {
		fmt.Fprintln(os.Stderr, "solve error:", err)
		os.Exit(1)
	}

	fmt.Println(total)
}

func readLines() ([]string, int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	maxLen := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > maxLen {
			maxLen = len(line)
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, 0, err
	}

	return lines, maxLen, nil
}

func padLines(lines []string, width int) [][]byte {
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		row := []byte(line)
		if len(row) < width {
			row = append(row, make([]byte, width-len(row))...)
		}
		grid[i] = row
	}
	return grid
}

func solveGrid(grid [][]byte, part2 bool) (int, error) {
	height := len(grid)
	if height == 0 {
		return 0, nil
	}

	width := len(grid[0])
	opRow := height - 1

	separator := make([]bool, width)
	hasSeparator := false
	for col := 0; col < width; col++ {
		allSpace := true
		for row := 0; row < height; row++ {
			if grid[row][col] != ' ' {
				allSpace = false
				break
			}
		}
		separator[col] = allSpace
		if allSpace {
			hasSeparator = true
		}
	}

	type problem struct {
		op      byte
		col     int
		numbers []int
	}

	ops := make([]problem, 0)
	for col, ch := range grid[opRow] {
		if ch == '+' || ch == '*' {
			ops = append(ops, problem{op: ch, col: col, numbers: make([]int, 0)})
		}
	}
	if len(ops) == 0 {
		return 0, fmt.Errorf("no operators found")
	}

	nearestOp := func(col int) int {
		bestIdx := 0
		bestDist := abs(col - ops[0].col)
		for i := 1; i < len(ops); i++ {
			dist := abs(col - ops[i].col)
			if dist < bestDist || (dist == bestDist && ops[i].col > ops[bestIdx].col) {
				bestDist = dist
				bestIdx = i
			}
		}
		return bestIdx
	}

	if hasSeparator {
		total := 0
		for col := 0; col < width; {
			for col < width && separator[col] {
				col++
			}
			start := col
			for col < width && !separator[col] {
				col++
			}
			if start == col {
				continue
			}
			var value int
			var err error
			if part2 {
				value, err = evaluateSegmentPart2(grid, start, col)
			} else {
				value, err = evaluateSegment(grid, start, col)
			}
			if err != nil {
				return 0, err
			}
			total += value
		}
		return total, nil
	}

	if part2 {
		// Each column encodes a single number (top to bottom); assign to nearest operator.
		for col := 0; col < width; col++ {
			var b strings.Builder
			for row := 0; row < opRow; row++ {
				ch := grid[row][col]
				if ch >= '0' && ch <= '9' {
					b.WriteByte(ch)
				}
			}
			if b.Len() == 0 {
				continue
			}
			value, err := strconv.Atoi(b.String())
			if err != nil {
				return 0, fmt.Errorf("column %d: %w", col, err)
			}
			idx := nearestOp(col)
			ops[idx].numbers = append(ops[idx].numbers, value)
		}
	} else {
		for row := 0; row < opRow; row++ {
			line := grid[row]
			col := 0
			for col < width {
				if line[col] < '0' || line[col] > '9' {
					col++
					continue
				}
				start := col
				for col < width && line[col] >= '0' && line[col] <= '9' {
					col++
				}
				raw := string(line[start:col])
				value, err := strconv.Atoi(strings.TrimSpace(raw))
				if err != nil {
					return 0, fmt.Errorf("row %d columns %d-%d: %w", row, start, col, err)
				}
				center := (start + col - 1) / 2
				idx := nearestOp(center)
				ops[idx].numbers = append(ops[idx].numbers, value)
			}
		}
	}

	total := 0
	for _, p := range ops {
		if len(p.numbers) == 0 {
			return 0, fmt.Errorf("no numbers found for operator at column %d", p.col)
		}
		switch p.op {
		case '+':
			sum := 0
			for _, n := range p.numbers {
				sum += n
			}
			total += sum
		case '*':
			product := 1
			for _, n := range p.numbers {
				product *= n
			}
			total += product
		default:
			return 0, fmt.Errorf("unhandled operator %q at column %d", p.op, p.col)
		}
	}

	return total, nil
}

func evaluateSegment(grid [][]byte, start, end int) (int, error) {
	height := len(grid)
	opRow := height - 1

	op := byte(0)
	for _, ch := range grid[opRow][start:end] {
		if ch != ' ' {
			op = ch
			break
		}
	}
	if op != '+' && op != '*' {
		return 0, fmt.Errorf("invalid operator in columns %d-%d", start, end)
	}

	numbers := make([]int, 0, height-1)
	for row := 0; row < opRow; row++ {
		raw := strings.TrimSpace(string(grid[row][start:end]))
		if raw == "" {
			continue
		}
		value, err := strconv.Atoi(raw)
		if err != nil {
			return 0, fmt.Errorf("row %d columns %d-%d: %w", row, start, end, err)
		}
		numbers = append(numbers, value)
	}

	if len(numbers) == 0 {
		return 0, fmt.Errorf("no numbers found in columns %d-%d", start, end)
	}

	switch op {
	case '+':
		sum := 0
		for _, n := range numbers {
			sum += n
		}
		return sum, nil
	case '*':
		product := 1
		for _, n := range numbers {
			product *= n
		}
		return product, nil
	}

	return 0, fmt.Errorf("unhandled operator %q", op)
}

func evaluateSegmentPart2(grid [][]byte, start, end int) (int, error) {
	height := len(grid)
	opRow := height - 1

	op := byte(0)
	for _, ch := range grid[opRow][start:end] {
		if ch != ' ' {
			op = ch
			break
		}
	}
	if op != '+' && op != '*' {
		return 0, fmt.Errorf("invalid operator in columns %d-%d", start, end)
	}

	numbers := make([]int, 0)
	for col := end - 1; col >= start; col-- {
		var b strings.Builder
		for row := 0; row < opRow; row++ {
			ch := grid[row][col]
			if ch >= '0' && ch <= '9' {
				b.WriteByte(ch)
			}
		}
		if b.Len() == 0 {
			continue
		}
		value, err := strconv.Atoi(b.String())
		if err != nil {
			return 0, fmt.Errorf("columns %d-%d: %w", start, end, err)
		}
		numbers = append(numbers, value)
	}

	if len(numbers) == 0 {
		return 0, fmt.Errorf("no numbers found in columns %d-%d", start, end)
	}

	switch op {
	case '+':
		sum := 0
		for _, n := range numbers {
			sum += n
		}
		return sum, nil
	case '*':
		product := 1
		for _, n := range numbers {
			product *= n
		}
		return product, nil
	}

	return 0, fmt.Errorf("unhandled operator %q", op)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
