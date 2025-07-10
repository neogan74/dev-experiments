package palindromenumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{
			name: "Test 1",
			x:    121,
			want: true,
		},
		{
			name: "Test 2",
			x:    -121,
			want: false,
		},
		{
			name: "Test 3",
			x:    10,
			want: false,
		},
	}

	for _, test := range tests {
		got := isPalindrome(test.x)
		if got != test.want {
			t.Errorf("Test %s failed, want %v got %v", test.name, test.want, got)
		}
	}
}
