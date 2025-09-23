package compareversionnumbers

import "testing"

func Test_compareVersion(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		version1 string
		version2 string
		want     int
	}{
		{
			name:     "Test Case 1",
			version1: "1.2",
			version2: "1.10",
			want:     -1,
		},
		{
			name:     "Test Case 2",
			version1: "1.01",
			version2: "1.001",
			want:     0,
		},
		{
			name:     "Test Case 3",
			version1: "1.0",
			version2: "1.0.0.0",
			want:     0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compareVersion(tt.version1, tt.version2)

			if got != tt.want {
				t.Errorf("compareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
