package taking_maximum_energy_from_the_mystic_dungeon

import "testing"

func Test_maxEnergy(t *testing.T) {
	type args struct {
		energy []int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				energy: []int{5, 2, -10, -5, 1},
				k:      3,
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				energy: []int{-2, -3, -1},
				k:      2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnergy(tt.args.energy, tt.args.k); got != tt.want {
				t.Errorf("maxEnergy() = %v, want %v", got, tt.want)
			}
		})
	}
}
