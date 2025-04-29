package _224_BasicCalculator

import (
	"fmt"
	"unicode"
)

func calculate(s string) int {
	stack := []int{}
	result, num, sign := 0, 0, 1

	for i := 0; i < len(s); i++ {
		char := s[i]

		switch {
		case unicode.IsDigit(rune(char)):
			num = num*10 + int(char-'0')

		case char == '+':
			result += sign * num
			num = 0
			sign = 1

		case char == '-':
			result += sign * num
			num = 0
			sign = -1

		case char == '(':
			// сохраняем текущий результат и знак в стек
			stack = append(stack, result, sign)
			// обнуляем текущий результат и знак для новой скобки
			result, sign = 0, 1

		case char == ')':
			result += sign * num
			num = 0
			// извлекаем предыдущий знак и результат
			sign = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			prevResult := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = prevResult + sign*result
		}
	}

	result += sign * num
	return result
}

func main() {
	fmt.Println(calculate("1 + 1"))               // 2
	fmt.Println(calculate(" 2-1 + 2 "))           // 3
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)")) // 23
	fmt.Println(calculate("1 + (2 - (3 + 4))"))   // -4
}
