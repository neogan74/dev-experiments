package main

import (
	"strings"
	"testing"
)

const exampleInput = `3-5
10-14
16-20
12-18

1
5
8
11
17
32
`

func TestExamplePart1(t *testing.T) {
	ranges, ids, err := loadInput(strings.NewReader(exampleInput))
	if err != nil {
		t.Fatalf("loadInput: %v", err)
	}
	if got, want := countFreshIDs(ids, ranges), 3; got != want {
		t.Fatalf("countFreshIDs = %d, want %d", got, want)
	}
}

func TestExamplePart2(t *testing.T) {
	ranges, _, err := loadInput(strings.NewReader(exampleInput))
	if err != nil {
		t.Fatalf("loadInput: %v", err)
	}
	if got, want := countFreshRangeIDs(ranges), 14; got != want {
		t.Fatalf("countFreshRangeIDs = %d, want %d", got, want)
	}
}

func TestParseRangeSwappedBounds(t *testing.T) {
	r, err := parseRange("20-10")
	if err != nil {
		t.Fatalf("parseRange returned error: %v", err)
	}
	if r.start != 10 || r.end != 20 {
		t.Fatalf("parseRange swapped bounds incorrectly: %+v", r)
	}
}
