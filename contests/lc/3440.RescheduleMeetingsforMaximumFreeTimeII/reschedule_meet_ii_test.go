package reschedulemeetingsformaximumfreetimeii

import "testing"

func TestMaxFreeTime(t *testing.T) {
	type args struct {
		eventTime int
		startTime []int
		endTime   []int
	}

	tests := []struct {
		name string
		args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				eventTime: 5,
				startTime: []int{1, 3},
				endTime:   []int{2, 5},
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				eventTime: 10,
				startTime: []int{0, 7, 9},
				endTime:   []int{1, 8, 10},
			},
			want: 7,
		},
		{
			name: "Test 3",
			args: args{
				eventTime: 5,
				startTime: []int{0, 1, 2, 3, 4},
				endTime:   []int{1, 2, 3, 4, 5},
			},
			want: 0,
		},
		{
			name: "Test 4",
			args: args{
				eventTime: 10,
				startTime: []int{0, 3, 7, 9},
				endTime:   []int{1, 4, 8, 10},
			},
			want: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := maxfreeTime(test.eventTime, test.startTime, test.endTime)
			if got != test.want {
				t.Errorf("maxFreeTime return %d, want %d", got, test.want)
			}
		})
	}
}
