package convertbinarynumberinalinkedlisttointeger

import "testing"

func Test_getDecimalValue(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Test 1",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
			wantAns: 5,
		},
		{
			name: "Test 2",
			args: args{
				head: &ListNode{
					Val: 0,
				},
			},
			wantAns: 0,
		},
		{
			name: "Test 3",
			args: args{
				head: &ListNode{
					Val: 1,
				},
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := getDecimalValue(tt.args.head); gotAns != tt.wantAns {
				t.Errorf("getDecimalValue() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
