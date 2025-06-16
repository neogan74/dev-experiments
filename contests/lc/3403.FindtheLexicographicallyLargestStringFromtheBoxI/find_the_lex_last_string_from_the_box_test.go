package _3403_find_the_lex_larg_string_from_the_box

import "testing"

func Test_answerString(t *testing.T) {
	type args struct {
		word       string
		numFriends int
	}
	tests := []struct {
		name    string
		args    args
		wantAns string
	}{
		{
			name: "Test 1",
			args: args{
				word:       "dbca",
				numFriends: 2,
			},
			wantAns: "dbc",
		},
		{
			name: "Test 2",
			args: args{
				word:       "gggg",
				numFriends: 4,
			},
			wantAns: "g",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := answerString(tt.args.word, tt.args.numFriends); gotAns != tt.wantAns {
				t.Errorf("answerString() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
