package maximumscorefromremovingsubstrings

import (
	"strings"
	"testing"
)

func Test_maximumGain(t *testing.T) {
	type args struct {
		s string
		x int
		y int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "test1",
			args: args{
				s: "cdbcbbaaabab",
				x: 4,
				y: 5,
			},
			wantAns: 19,
		},
		{
			name: "test2",
			args: args{
				s: "aabbaaxybbaabb",
				x: 5,
				y: 4,
			},
			wantAns: 20,
		},
		{
			name: "empty string",
			args: args{
				s: "",
				x: 10,
				y: 10,
			},
			wantAns: 0,
		},
		{
			name: "no ab or ba",
			args: args{
				s: "cccccccc",
				x: 3,
				y: 2,
			},
			wantAns: 0,
		},
		{
			name: "all ab",
			args: args{
				s: "ababab",
				x: 2,
				y: 1,
			},
			wantAns: 6,
		},
		{
			name: "all ba",
			args: args{
				s: "bababa",
				x: 1,
				y: 2,
			},
			wantAns: 6,
		},
		{
			name: "mixed with other chars",
			args: args{
				s: "abxybazb",
				x: 3,
				y: 2,
			},
			wantAns: 5,
		},
		{
			name: "long string",
			args: args{
				s: strings.Repeat("ab", 1000) + strings.Repeat("ba", 1000),
				x: 2,
				y: 3,
			},
			wantAns: 5999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximumGain(tt.args.s, tt.args.x, tt.args.y); gotAns != tt.wantAns {
				t.Errorf("maximumGain() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}

	// Benchmark
}

func Benchmark_maximumGain(b *testing.B) {
	cases := []struct {
		name string
		s    string
		x    int
	}{
		{"medium", strings.Repeat("ab", 1000) + strings.Repeat("ba", 1000), 3},
		{"long", strings.Repeat("ab", 10000) + strings.Repeat("ba", 10000), 5},
		{"short", "ababa", 2},
	}
	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				maximumGain(c.s, c.x, 2) // Using a fixed y value for benchmarking
			}
		})
	}
}
