package main

import "testing"

const exampleInput = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

func TestSolveLinePart1Example(t *testing.T) {
	got := solveLine(exampleInput, false)
	const want int64 = 1227775554
	if got != want {
		t.Fatalf("part1 example: got %d, want %d", got, want)
	}
}

func TestSolveLinePart2Example(t *testing.T) {
	got := solveLine(exampleInput, true)
	const want int64 = 4174379265
	if got != want {
		t.Fatalf("part2 example: got %d, want %d", got, want)
	}
}

func TestSmallRanges(t *testing.T) {
	cases := []struct {
		line    string
		partTwo bool
		want    int64
	}{
		{"11-22", false, 33}, // 11 + 22
		{"11-22", true, 33},  // same for part 2
		{"95-115", false, 99},
		{"95-115", true, 99 + 111},
	}
	for _, tc := range cases {
		got := solveLine(tc.line, tc.partTwo)
		if got != tc.want {
			t.Fatalf("line %q part2=%v: got %d, want %d", tc.line, tc.partTwo, got, tc.want)
		}
	}
}
