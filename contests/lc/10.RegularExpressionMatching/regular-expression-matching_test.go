package regularexpressionmatching

import "testing"

func Test_isMatch(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want bool
	}{
		{
			name: "Example 1",
			s:    "aa",
			p:    "a",
			want: false,
		},
		{
			name: "Example 2",
			s:    "aa",
			p:    "a*",
			want: true,
		},
		{
			name: "Example 3",
			s:    "ab",
			p:    ".*",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isMatch(tt.s, tt.p)
			if got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsMatch(b *testing.B) {
	// Примеры входных данных для тестирования
	testCases := []struct {
		s, p string
	}{
		{"aa", "a*"},
		{"ab", ".*"},
		{"aab", "c*a*b"},
		// Добавьте дополнительные тест-кейсы по мере необходимости
	}

	// Запуск бенчмарка
	for _, tc := range testCases {
		b.Run(tc.s+"_"+tc.p, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isMatch(tc.s, tc.p)
			}
		})
	}
}
