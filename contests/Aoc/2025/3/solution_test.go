package main

import "testing"

func TestMaxBankJoltageExamples(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		pick     int
		expected int64
	}{
		{name: "part1_first_bank", line: "987654321111111", pick: 2, expected: 98},
		{name: "part1_second_bank", line: "811111111111119", pick: 2, expected: 89},
		{name: "part1_third_bank", line: "234234234234278", pick: 2, expected: 78},
		{name: "part1_fourth_bank", line: "818181911112111", pick: 2, expected: 92},
		{name: "part2_first_bank", line: "987654321111111", pick: 12, expected: 987654321111},
		{name: "part2_second_bank", line: "811111111111119", pick: 12, expected: 811111111119},
		{name: "part2_third_bank", line: "234234234234278", pick: 12, expected: 434234234278},
		{name: "part2_fourth_bank", line: "818181911112111", pick: 12, expected: 888911112111},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := maxBankJoltage(tt.line, tt.pick)
			if got != tt.expected {
				t.Fatalf("maxBankJoltage(%q, %d) = %d, want %d", tt.line, tt.pick, got, tt.expected)
			}
		})
	}
}

func TestTotalsFromExample(t *testing.T) {
	lines := []string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	}

	tests := []struct {
		name     string
		pick     int
		expected int64
	}{
		{name: "part1_total", pick: 2, expected: 357},
		{name: "part2_total", pick: 12, expected: 3121910778619},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			var total int64
			for _, line := range lines {
				total += maxBankJoltage(line, tt.pick)
			}
			if total != tt.expected {
				t.Fatalf("total = %d, want %d", total, tt.expected)
			}
		})
	}
}
