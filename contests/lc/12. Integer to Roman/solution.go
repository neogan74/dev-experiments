package main

import (
	"fmt"
)

func intToRoman(num int) string {
	romanNumerals := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	var result string
	for _, value := range []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1} {
		for num >= value {
			result += romanNumerals[value]
			num -= value
		}
	}

	return result
}

func main() {
	fmt.Println(intToRoman(3749)) // Output: "MMMDCCXLIX"
	fmt.Println(intToRoman(58))   // Output: "LVIII"
	fmt.Println(intToRoman(1994)) // Output: "MCMXCIV"
}
