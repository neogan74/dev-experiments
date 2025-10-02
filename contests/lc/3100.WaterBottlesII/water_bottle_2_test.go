package _100_WaterBottlesII

import "testing"

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
