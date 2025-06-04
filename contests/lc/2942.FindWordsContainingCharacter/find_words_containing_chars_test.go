package _2942_find_words_conitaining_character

import (
	"reflect"
	"testing"
)

func Test_findWordsContaining(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		x     byte
		want  []int
	}{
		{
			name:  "Test 1",
			words: []string{"leet", "code"},
			x:     'e',
			want:  []int{0, 1},
		},
		{
			name:  "Test 2",
			words: []string{"abc", "bcd", "aaaa", "cbc"},
			x:     'a',
			want:  []int{0, 2},
		},
		// {
		// 	name:  "Test 3",
		// 	words: []string{"abc", "bcd", "aaaa", "cbc"},
		// 	x:     'z',
		// 	want:  []int{},
		// },
	}
	for _, tt := range tests {
		got := findWordsContaining(tt.words, tt.x)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findWordsContaining() = %v, want %v", got, tt.want)
		}
	}
}
