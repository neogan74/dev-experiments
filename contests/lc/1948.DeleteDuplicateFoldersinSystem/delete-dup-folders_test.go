package deleteduplicatefolders

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicateFolder(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// {
		// 	name: "Test 1",
		// 	args: args{
		// 		paths: [][]string{{"a"}, {"b"}, {"c"}, {"d"}, {"a", "b"}, {"c", "b"}, {"d", "a"}},
		// 	},
		// 	want: [][]string{{"b"}, {"d"}, {"d", "a"}},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicateFolder(tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicateFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}
