package _1128_number_of_equivalent_domino_pairs

func numEquivDominoPairs(dominoes [][]int) int {
	count := make([]int, 100)
	res := 0
	for _, d := range dominoes {
		val := d[0]*10 + d[1]
		if d[0] > d[1] {
			val = d[1]*10 + d[0]
		}
		res += count[val]
		count[val]++
	}
	return res
}
