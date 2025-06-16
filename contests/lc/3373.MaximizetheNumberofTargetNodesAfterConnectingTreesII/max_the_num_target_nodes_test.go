package _3373_max_the_num_of_target_nodes_after

import (
	"reflect"
	"testing"
)

func Test_maxTargetNodes(t *testing.T) {
	type args struct {
		edges1 [][]int
		edges2 [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		//TODO:
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTargetNodes(tt.args.edges1, tt.args.edges2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxTargetNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
