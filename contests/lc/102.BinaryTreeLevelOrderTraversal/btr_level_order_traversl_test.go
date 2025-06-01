package _102_binary_tree_level_order_traversal

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "Test 1",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   9,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 20,
						Left: &TreeNode{
							Val:   15,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			wantAns: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "Test 2",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
			wantAns: [][]int{{1}},
		},
		// {
		// 	name: "Test 3",
		// 	args: args{
		// 		root: nil,
		// 	},
		// 	wantAns: [][]int{},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := levelOrder(tt.args.root); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("levelOrder() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
