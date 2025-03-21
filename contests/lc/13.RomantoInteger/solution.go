package main

import "fmt"

func romanToInt(s string) int {
	symbols := map[byte]int{
		'I': 1, 'V': 5, 'X': 10,
		'L': 50, 'C': 100,
		'D': 500, 'M': 1000,
	}

	res := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && symbols[s[i]] < symbols[s[i+1]] {
			res -= symbols[s[i]]
		} else {
			res += symbols[s[i]]
		}
	}
	return res
}

func main() {
	fmt.Println(romanToInt("III"))     // 3
	fmt.Println(romanToInt("LVIII"))   // 58
	fmt.Println(romanToInt("MCMXCIV")) // 1994
}
