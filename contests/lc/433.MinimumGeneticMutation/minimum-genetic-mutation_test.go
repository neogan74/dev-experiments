package minimumgeneticmutation

import "testing"

func Test_minMutation(t *testing.T) {
	type args struct {
		start string
		end   string
		bank  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				start: "AACCGGTT",
				end:   "AACCGGTA",
				bank:  []string{"AACCGGTA"},
			},
			want: 1,
		},
		{
			name: "Test 1",
			args: args{
				start: "AACCGGTT",
				end:   "AAACGGTA",
				bank:  []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMutation(tt.args.start, tt.args.end, tt.args.bank); got != tt.want {
				t.Errorf("minMutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
