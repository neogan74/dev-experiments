package _138_copy_list_with_random_pointers

import (
	"reflect"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	//test 1
	node1 := Node{}

	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "test1",
			args: args{
				head: &node1,
			},
			want: &node1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyRandomList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyRandomList() = %v, want %v", got, tt.want)
			}
		})
	}
}
