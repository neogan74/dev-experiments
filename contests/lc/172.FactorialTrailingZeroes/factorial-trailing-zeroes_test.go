package factorialtrailingzeroes

import "testing"

func Test_trailingZeroes(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "test case 1",
			n:    3,
			want: 0,
		},
		{
			name: "test case 2",
			n:    5,
			want: 1,
		},
		{
			name: "test case 3",
			n:    10,
			want: 2,
		},
		{
			name: "test case 4",
			n:    0,
			want: 0,
		},
		{
			name: "test case 5",
			n:    10000,
			want: 2499,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroes(tt.n); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Бенчмарки
func BenchmarkTrailingZeroesSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trailingZeroes(5)
	}
}

func BenchmarkTrailingZeroesMedium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trailingZeroes(1000)
	}
}

func BenchmarkTrailingZeroesLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trailingZeroes(1_000_000)
	}
}
