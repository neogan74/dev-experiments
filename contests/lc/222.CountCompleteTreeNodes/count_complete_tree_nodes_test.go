package _222_count_complete_tree_nodes

import "testing"

func Test_countNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
						Right: &TreeNode{Val: 5, Left: nil, Right: nil}},
					Right: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
						Right: nil,
					},
				},
			},
			want: 6,
		},
		{
			name: "Test 2",
			args: args{
				root: nil,
			},
			want: 0,
		},
		{
			name: "Test 3",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodes(tt.args.root); got != tt.want {
				t.Errorf("countNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
