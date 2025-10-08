package _488_AvoidFloodinTheCity

import "sort"

func AvoidFlood(rains []int) []int {
	n := len(rains)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = 1
	}

	last := make(map[int]int) // lake -> last rainy day
	dry := make([]int, 0, n)  // sorted indices of dry days

	for i, x := range rains {
		if x == 0 {
			// запомним сухой день
			dry = append(dry, i)
		} else {
			ans[i] = -1
			if prev, ok := last[x]; ok {
				// ищем первый dry day > prev
				pos := sort.SearchInts(dry, prev+1)
				if pos == len(dry) {
					return []int{} // невозможно
				}
				j := dry[pos]
				ans[j] = x
				// удалить dry[pos]
				dry = append(dry[:pos], dry[pos+1:]...)
			}
			last[x] = i
		}
	}
	return ans
}
