package findtheoriginaltypedstringii

import "testing"

// Тесты для ModInt
func TestModIntBasicOperations(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected struct {
			add ModInt
			sub ModInt
			mul ModInt
		}
	}{
		{
			name: "Small numbers",
			a:    5, b: 3,
			expected: struct {
				add ModInt
				sub ModInt
				mul ModInt
			}{
				add: ModInt(8),
				sub: ModInt(2),
				mul: ModInt(15),
			},
		},
		{
			name: "Large numbers",
			a:    500000000, b: 700000000,
			expected: struct {
				add ModInt
				sub ModInt
				mul ModInt
			}{
				add: ModInt((500000000 + 700000000) % MOD),
				sub: ModInt((500000000 - 700000000 + MOD) % MOD),
				mul: ModInt((uint64(500000000) * uint64(700000000)) % MOD),
			},
		},
		{
			name: "Overflow add",
			a:    MOD - 1, b: 2,
			expected: struct {
				add ModInt
				sub ModInt
				mul ModInt
			}{
				add: ModInt(1),
				sub: ModInt(MOD - 3),
				mul: ModInt((uint64(MOD-1) * uint64(2)) % MOD),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, b := NewModInt(tt.a), NewModInt(tt.b)
			add := a.Add(b)
			if add != tt.expected.add {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, add, tt.expected.add)
			}

			sub := a.Sub(b)
			if sub != tt.expected.sub {
				t.Errorf("Sub(%d, %d) = %d, want %d", tt.a, tt.b, sub, tt.expected.sub)
			}
			mul := a.Mul(b)
			if mul != tt.expected.mul {
				t.Errorf("Mul(%d, %d) = %d, want %d", tt.a, tt.b, mul, tt.expected.mul)
			}
		})
	}
}

func TestModIntPow(t *testing.T) {
	tests := []struct {
		base int
		exp  int
		want ModInt
	}{
		{2, 10, ModInt(1024)},
		{3, 4, ModInt(81)},
		{5, 0, ModInt(1)},
		{7, 1, ModInt(7)},
		{10, 9, ModInt(1000000000)},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			base := NewModInt(tt.base)
			got := base.Pow(tt.exp)
			if got != tt.want {
				t.Errorf("ModInt(%d).Pow(%d) = %d, want %d", tt.base, tt.exp, got, tt.want)
			}
		})
	}
}

func TestModIntInvAndDiv(t *testing.T) {
	tests := []struct {
		a, b int
		want ModInt
	}{
		{10, 2, ModInt(5)},
		{7, 3, ModInt(7 * ModInt(3).Pow(MOD-2).Mul(1))},
		{1, 1, ModInt(1)},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			a, b := NewModInt(tt.a), NewModInt(tt.b)
			got := a.Div(b)
			if got != tt.want {
				t.Errorf("ModInt(%d).Div(%d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}

			// Проверка, что a/b * b = a
			check := got.Mul(b)
			if check != a {
				t.Errorf("ModInt(%d).Div(%d).Mul(%d) = %d, want %d", tt.a, tt.b, tt.b, check, a)
			}
		})
	}
}

// Тесты для possibleStringCount
func TestPossibleStringCount(t *testing.T) {
	tests := []struct {
		name string
		word string
		k    int
		want int
	}{
		{
			name: "Simple case",
			word: "aabbc",
			k:    3,
			want: 4, // 2*2*1 = 4 (возможные строки с 3 группами)
		},
		{
			name: "Single character",
			word: "a",
			k:    1,
			want: 1,
		},
		{
			name: "Repeated characters",
			word: "aaa",
			k:    1,
			want: 3,
		},
		{
			name: "Two groups",
			word: "aaabbb",
			k:    2,
			want: 9, // 3*3 = 9
		},
		{
			name: "k greater than groups",
			word: "aabb",
			k:    3,
			want: 3, // невозможно получить 3 группы из 2 групп
		},
		{
			name: "Complex case",
			word: "aaabbbccc",
			k:    2,
			want: 27, // 3*3*3 = 27 (возможные строки с 2 группами)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := possibleStringCount(tt.word, tt.k)
			if got != tt.want {
				t.Errorf("possibleStringCount(%q, %d) = %d, want %d", tt.word, tt.k, got, tt.want)
			}
		})
	}
}

// Бенчмарки для оценки производительности
func BenchmarkPossibleStringCount(b *testing.B) {
	testCases := []struct {
		word string
		k    int
	}{
		{"aabbc", 3},
		{"aaabbbccc", 2},
		{"aaaabbbbccccdddd", 4},
	}

	for _, tc := range testCases {
		b.Run("", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				possibleStringCount(tc.word, tc.k)
			}
		})
	}
}
