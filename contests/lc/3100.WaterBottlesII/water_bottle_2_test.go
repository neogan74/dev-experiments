package _100_WaterBottlesII

import (
	"fmt"
	"testing"
)

func Test_maxBottlesDrunk(t *testing.T) {
	type args struct {
		numBottles  int
		numExchange int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				numBottles:  13,
				numExchange: 6,
			},
			want: 15,
		},
		{
			name: "Test 2",
			args: args{
				numBottles:  10,
				numExchange: 3,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxBottlesDrunk(tt.args.numBottles, tt.args.numExchange); got != tt.want {
				t.Errorf("maxBottlesDrunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maxBottlesDrunk(b *testing.B) {
	cases := []struct {
		name        string
		numBottles  int
		numExchange int
	}{
		{name: "SmallExchange", numBottles: 10, numExchange: 3},
		{name: "MidExchange", numBottles: 40, numExchange: 6},
		{name: "LargeExchange", numBottles: 100, numExchange: 12},
	}

	for _, tc := range cases {
		b.Run(fmt.Sprintf("%s_%d_%d", tc.name, tc.numBottles, tc.numExchange), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				maxBottlesDrunk(tc.numBottles, tc.numExchange)
			}
		})
	}
}
