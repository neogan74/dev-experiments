package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	in.Buffer(make([]byte, 0, 1<<20), 1<<20)

	nextInt := func() int {
		if !in.Scan() {
			return 0
		}
		v, _ := strconv.Atoi(in.Text())
		return v
	}

	n := nextInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = nextInt()
	}
	x := nextInt()

	best := arr[0]
	minDiff := abs(arr[0] - x)
	for i := 1; i < n; i++ {
		d := abs(arr[i] - x)
		if d < minDiff {
			minDiff = d
			best = arr[i]
		}
		// если d == minDiff — ничего не делаем, по условию можно вывести любой
	}

	out := bufio.NewWriter(os.Stdout)
	_, err := out.WriteString(strconv.Itoa(best))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
	err = out.WriteByte('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
	out.Flush()
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
