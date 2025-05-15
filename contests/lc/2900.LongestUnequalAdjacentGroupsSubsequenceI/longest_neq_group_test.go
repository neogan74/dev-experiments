package _2900_longest_neq_group

import (
	"reflect"
	"testing"
)

func Test_getLongestSubsequence(t *testing.T) {
	type args struct {
		words  []string
		groups []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name: "Test 1",
			args: args{
				words:  []string{"c"},
				groups: []int{0},
			},
			wantAns: []string{"c"},
		},
		{
			name: "Test 1",
			args: args{
				words:  []string{"d"},
				groups: []int{1},
			},
			wantAns: []string{"d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := getLongestSubsequence(tt.args.words, tt.args.groups); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("getLongestSubsequence() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
