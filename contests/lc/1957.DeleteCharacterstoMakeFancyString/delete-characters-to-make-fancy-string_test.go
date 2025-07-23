package deletecharacterstomakefancystring

import "testing"

func TestMakeFancyString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},                // Пустая строка
		{"a", "a"},              // Одна буква
		{"aa", "aa"},            // Две одинаковые буквы
		{"aaa", "aa"},           // Три одинаковые буквы
		{"aaaa", "aa"},          // Четыре одинаковые буквы
		{"abba", "abba"},        // Нормальная строка без лишних повторов
		{"aaabbbccc", "aabbcc"}, // Много повторений подряд
		{"abcabc", "abcabc"},    // Нет повторяющихся соседей
		{"aabbaaa", "aabbaa"},   // Смешанные повторы
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := makeFancyString(test.input)
			if result != test.expected {
				t.Errorf("makeFancyString(%q) = %q, want %q", test.input, result, test.expected)
			}
		})
	}
}
