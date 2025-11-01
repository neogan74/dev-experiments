package _115_PrintFooBarAlternately

import (
	"reflect"
	"testing"
)

func TestNewFooBar(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want *FooBar
	}{
		{
			name: "Test 1",
			args: args{
				n: 1,
			},
			want: &FooBar{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFooBar(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFooBar() = %v, want %v", got, tt.want)
			}
		})
	}
}
