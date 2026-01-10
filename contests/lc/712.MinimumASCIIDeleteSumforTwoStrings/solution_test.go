package main

import "testing"

func TestMinimumDeleteSum(t *testing.T) {
	cases := []struct {
		s1   string
		s2   string
		want int
	}{
		{s1: "sea", s2: "eat", want: 231},
		{s1: "delete", s2: "leet", want: 403},
		{s1: "", s2: "abc", want: int('a' + 'b' + 'c')},
		{s1: "abc", s2: "", want: int('a' + 'b' + 'c')},
		{s1: "a", s2: "a", want: 0},
		{s1: "a", s2: "b", want: int('a' + 'b')},
		{s1: "ab", s2: "ba", want: int('a' + 'b')},
	}

	for _, tc := range cases {
		if got := minimumDeleteSum(tc.s1, tc.s2); got != tc.want {
			t.Fatalf("minimumDeleteSum(%q, %q) = %d, want %d", tc.s1, tc.s2, got, tc.want)
		}
	}
}
