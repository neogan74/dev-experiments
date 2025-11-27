package _3228_MaximumNumberofOperationstoMoveOnestotheEnd

import (
	"strings"
	"testing"
)

func TestKnownCases(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want int
	}{
		{"example1", "1001101", 4},
		{"example2", "00111", 0},
		{"singleOne", "1", 0},
		{"trailingZero", "10", 1},
		{"mixed", "010010", 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := maxOperations(tc.s); got != tc.want {
				t.Fatalf("maxOperations(%q) = %d, want %d", tc.s, got, tc.want)
			}
		})
	}
}

func TestMatchesBruteForceUpToSixBits(t *testing.T) {
	memo := map[string]int{}
	var dfs func(string) int
	dfs = func(state string) int {
		if v, ok := memo[state]; ok {
			return v
		}
		best := 0
		for _, next := range enumerateOperations(state) {
			if candidate := 1 + dfs(next); candidate > best {
				best = candidate
			}
		}
		memo[state] = best
		return best
	}

	for n := 1; n <= 6; n++ {
		limit := 1 << n
		for mask := 0; mask < limit; mask++ {
			var sb strings.Builder
			for bit := n - 1; bit >= 0; bit-- {
				if mask&(1<<bit) != 0 {
					sb.WriteByte('1')
				} else {
					sb.WriteByte('0')
				}
			}
			s := sb.String()
			if got, want := maxOperations(s), dfs(s); got != want {
				t.Fatalf("maxOperations(%q) = %d, want %d", s, got, want)
			}
		}
	}
}

func enumerateOperations(state string) []string {
	arr := []byte(state)
	n := len(arr)
	var out []string

	for i := 0; i < n-1; i++ {
		if arr[i] == '1' && arr[i+1] == '0' {
			next := append([]byte{}, arr...)
			j := i
			for j+1 < n && next[j+1] == '0' {
				next[j], next[j+1] = next[j+1], next[j]
				j++
			}
			out = append(out, string(next))
		}
	}

	return out
}
