package _3343_count_number_of_balanced_prenumtations

import "testing"

func Test_countBalancedPremutations(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBalancedPremutations(tt.args.num); got != tt.want {
				t.Errorf("countBalancedPremutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
