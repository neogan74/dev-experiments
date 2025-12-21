package _55_DeleteColumnstoMakeSortedII

import "testing"

func Test_minDeletionSize(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				strs: []string{"ca", "bb", "ac"},
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				strs: []string{"xc", "yb", "za"},
			},
			want: 0,
		},
		{
			name: "Test 3",
			args: args{
				strs: []string{"xyz", "wvu", "tsr"},
			},
			want: 3,
		},
		{
			name: "Test 4",
			args: args{
				strs: []string{"a", "b", "c"},
			},
			want: 0,
		},
		{
			name: "Test 5",
			args: args{
				strs: []string{"ab", "aa"},
			},
			want: 1,
		},
		{
			name: "Test 6",
			args: args{
				strs: []string{"aa", "aa", "aa"},
			},
			want: 0,
		},
		{
			name: "Test 7",
			args: args{
				strs: []string{"xga", "xfb", "yfa"},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDeletionSize(tt.args.strs); got != tt.want {
				t.Errorf("minDeletionSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
