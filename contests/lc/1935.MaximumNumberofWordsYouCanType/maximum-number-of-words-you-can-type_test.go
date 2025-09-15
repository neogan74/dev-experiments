package maximumnumberofwordsyoucantype

import "testing"

func Test_canBeTypedWords(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		text          string
		brokenLetters string
		want          int
	}{
		{
			name:          "Example 1",
			text:          "hello world",
			brokenLetters: "ad",
			want:          1,
		},
		{
			name:          "Example 2",
			text:          "leet code",
			brokenLetters: "lt",
			want:          1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canBeTypedWords(tt.text, tt.brokenLetters)
			if got != tt.want {
				t.Errorf("canBeTypedWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
