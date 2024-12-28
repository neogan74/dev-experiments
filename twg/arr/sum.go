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

func SumAllTails(numToTailSum ...[]int) []int {
	var sums []int
	for _, numbers := range numToTailSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
