package main

import (
	"math/big"
	"strconv"
)

// Функция для подсчёта факториала
func factorial(n int) *big.Int {
	res := big.NewInt(1)
	for i := 2; i <= n; i++ {
		res.Mul(res, big.NewInt(int64(i)))
	}
	return res
}

// Функция для подсчёта количества уникальных перестановок цифр
func countPermutations(digitCount map[int]int, totalDigits int) *big.Int {
	res := factorial(totalDigits)
	for _, count := range digitCount {
		if count > 1 {
			res.Div(res, factorial(count))
		}
	}
	return res
}

// Основная функция для подсчёта количества хороших чисел
func countGoodIntegers(n int, k int) int64 {
	if n == 1 {
		count := 0
		for i := 1; i <= 9; i++ {
			if i%k == 0 {
				count++
			}
		}
		return int64(count)
	}

	halfLength := (n + 1) / 2
	minHalf := intPow(10, halfLength-1)
	maxHalf := intPow(10, halfLength)
	seen := make(map[string]bool)
	total := big.NewInt(0)

	for num := minHalf; num < maxHalf; num++ {
		firstHalf := strconv.Itoa(num)
		var palindrome string
		if n%2 == 0 {
			palindrome = firstHalf + reverseString(firstHalf)
		} else {
			palindrome = firstHalf + reverseString(firstHalf[:len(firstHalf)-1])
		}

		palNum, _ := strconv.Atoi(palindrome)
		if palNum%k != 0 {
			continue
		}

		digits := []rune(palindrome)
		digitCount := make(map[int]int)
		for _, d := range digits {
			digitCount[int(d-'0')]++
		}

		// Проверка на ведущий ноль
		if digitCount[0] == n {
			continue
		}

		// Создание ключа для уникальности
		key := make([]rune, 0, len(digits))
		for d := 0; d <= 9; d++ {
			for i := 0; i < digitCount[d]; i++ {
				key = append(key, rune('0'+d))
			}
		}
		keyStr := string(key)
		if seen[keyStr] {
			continue
		}
		seen[keyStr] = true

		// Подсчёт перестановок
		perms := countPermutations(digitCount, n)

		// Вычитание перестановок с ведущим нулём
		if digitCount[0] > 0 {
			digitCount[0]--
			permsZero := countPermutations(digitCount, n-1)
			perms.Sub(perms, permsZero)
			digitCount[0]++
		}

		total.Add(total, perms)
	}

	return total.Int64()
}

// Вспомогательная функция для возведения в степень
func intPow(a, b int) int {
	res := 1
	for b > 0 {
		if b%2 == 1 {
			res *= a
		}
		a *= a
		b /= 2
	}
	return res
}

// Вспомогательная функция для разворота строки
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
