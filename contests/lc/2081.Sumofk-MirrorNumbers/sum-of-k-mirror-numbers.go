package sumofkmirrornumbers

import "strconv"

func isPalBase(x int, base int) bool {
	if x == 0 {
		return true
	}
	var digits []int
	for v := x; v > 0; v /= base {
		digits = append(digits, v%base)
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}
	return true
}

// Генерация k‑mirror чисел
func kMirror(k int, n int) int64 {
	found, sum := 0, int64(0)
	// длина сгенерированного палиндрома в decimal
	for length := 1; found < n; length++ {
		halfLen := (length + 1) / 2
		start := 1
		for i := 1; i < halfLen; i++ {
			start *= 10
		}
		end := start * 10

		for prefix := start; prefix < end && found < n; prefix++ {
			s := strconv.Itoa(prefix)
			rs := reverseStr(s[:len(s)-(length%2)])
			palStr := s + rs
			palNum, _ := strconv.Atoi(palStr)

			if isPalBase(palNum, k) {
				sum += int64(palNum)
				found++
			}
		}
	}
	return sum
}

// Помощь: реверсия строки
func reverseStr(s string) string {
	bs := []byte(s)
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}
