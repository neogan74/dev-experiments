package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputPath := flag.String("input", "input.txt", "path to input file")
	partTwo := flag.Bool("part2", false, "use part two rules (repeat >= 2 times)")
	flag.Parse()

	// Allow a positional override for the input path to stay convenient.
	if flag.NArg() > 0 {
		*inputPath = flag.Arg(0)
	}

	data, err := os.ReadFile(*inputPath)
	if err != nil {
		log.Fatalf("read %s: %v", *inputPath, err)
	}

	line := strings.TrimSpace(string(data))
	if line == "" {
		fmt.Println("0")
		return
	}

	fmt.Println(solveLine(line, *partTwo))
}

func solveLine(line string, partTwo bool) int64 {
	ranges, maxEnd := parseRanges(line)
	maxRepeat := 2
	if partTwo {
		maxRepeat = 0 // 0 means unlimited (>=2)
	}

	repeats := generateRepeatedNumbers(maxEnd, maxRepeat)
	prefix := make([]int64, len(repeats)+1)
	for i, v := range repeats {
		prefix[i+1] = prefix[i] + v
	}

	var total int64
	for _, rg := range ranges {
		l := sort.Search(len(repeats), func(i int) bool { return repeats[i] >= rg.start })
		r := sort.Search(len(repeats), func(i int) bool { return repeats[i] > rg.end })
		total += prefix[r] - prefix[l]
	}
	return total
}

func pow10Safe(n int) (int64, bool) {
	var r int64 = 1
	for i := 0; i < n; i++ {
		if r > math.MaxInt64/10 {
			return 0, false
		}
		r *= 10
	}
	return r, true
}

func powInt(base int64, exp int) (int64, bool) {
	r := int64(1)
	for i := 0; i < exp; i++ {
		if willOverflowMul(r, base) {
			return 0, false
		}
		r *= base
	}
	return r, true
}

func willOverflowMul(a, b int64) bool {
	if a == 0 || b == 0 {
		return false
	}
	if a == 1 || b == 1 {
		return false
	}
	if a == -1 {
		return b == math.MinInt64
	}
	if b == -1 {
		return a == math.MinInt64
	}
	// Ensure |a| <= |b|
	if a < 0 {
		if b < 0 {
			if a < b {
				a, b = b, a
			}
		} else {
			if -a > b {
				a, b = b, a
			}
		}
	} else if b < 0 && a > -b {
		a, b = b, a
	}

	return a > math.MaxInt64/b || a < math.MinInt64/b
}

func ceilDiv(a, b int64) int64 {
	if b == 0 {
		return 0
	}
	q := a / b
	if a%b != 0 {
		q++
	}
	return q
}

type numRange struct {
	start int64
	end   int64
}

func parseRanges(line string) ([]numRange, int64) {
	var rs []numRange
	var maxEnd int64
	for _, part := range strings.Split(line, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		bounds := strings.Split(part, "-")
		if len(bounds) != 2 {
			log.Fatalf("invalid range %q", part)
		}
		start, err := strconv.ParseInt(bounds[0], 10, 64)
		if err != nil {
			log.Fatalf("invalid start %q: %v", bounds[0], err)
		}
		end, err := strconv.ParseInt(bounds[1], 10, 64)
		if err != nil {
			log.Fatalf("invalid end %q: %v", bounds[1], err)
		}
		if start > end {
			start, end = end, start
		}
		if end > maxEnd {
			maxEnd = end
		}
		rs = append(rs, numRange{start: start, end: end})
	}
	return rs, maxEnd
}

func generateRepeatedNumbers(limit int64, maxRepeat int) []int64 {
	set := make(map[int64]struct{})
	for lenSeq := 1; ; lenSeq++ {
		pow, ok := pow10Safe(lenSeq)
		if !ok {
			break
		}
		minSeq := pow / 10 // smallest sequence without leading zero
		if minSeq == 0 {
			continue
		}
		maxSeq := pow - 1

		for repeat := 2; ; repeat++ {
			if maxRepeat > 0 && repeat > maxRepeat {
				break
			}

			powK, ok := powInt(pow, repeat)
			if !ok {
				break
			}
			geom := (powK - 1) / (pow - 1) // 1 + pow + pow^2 + ... + pow^(repeat-1)
			minVal := minSeq * geom
			if minVal < 0 || minVal > limit {
				break // further repeats will only get larger
			}
			endSeq := maxSeq
			maxVal := maxSeq * geom
			if maxVal < 0 {
				break
			}
			if maxVal > limit {
				endSeq = limit / geom
			}
			for seq := minSeq; seq <= endSeq; seq++ {
				val := seq * geom
				if val <= limit {
					set[val] = struct{}{}
				}
			}
		}
	}

	nums := make([]int64, 0, len(set))
	for v := range set {
		nums = append(nums, v)
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	return nums
}
