package _230_kth_smallest_element_in_bst

import "testing"

func Test_kthSmallest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
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
					Val: 3,
					Left: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val: 2,
						},
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				k: 1,
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val: 1,
							},
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				k: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
