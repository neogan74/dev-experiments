package _02_HappyNumber

// sumOfSquares возвращает сумму квадратов цифр n.
func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		d := n % 10
		sum += d * d
		n /= 10
	}
	return sum
}

// isHappy проверяет, является ли число счастливым.
func isHappy(n int) bool {
	seen := make(map[int]struct{})
	for n != 1 {
		if _, exists := seen[n]; exists {
			// зашли в цикл, не дойдя до 1
			return false
		}
		seen[n] = struct{}{}
		n = sumOfSquares(n)
	}
	return true
}
