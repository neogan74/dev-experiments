package arr

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numb := range numbersToSum {
		sums = append(sums, Sum(numb))
	}
	return sums
}
