package _2894_DivisibleandNon_divisibleSumsDifference

func differenceOfSums(n int64, m int64) int64 {
	total := n * (n + 1) / 2
	k := n / m
	sumDiv := m * (k * (k + 1) / 2)
	return total - 2*sumDiv
}
