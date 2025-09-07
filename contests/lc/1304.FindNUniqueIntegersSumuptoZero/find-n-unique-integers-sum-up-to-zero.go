package findnuniqueintegerssumuptozero

func sumZero(n int) []int {
	res := make([]int, 0, n)
	for i := 1; i <= n/2; i++ {
		res = append(res, i, -i)
	}
	if n%2 == 1 {
		res = append(res, 0)
	}
	return res
}
