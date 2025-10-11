package _186_MaximumTotalDamageWithSpellCasting

import "sort"

func MaximumTotalDamage(power []int) int64 {
	if len(power) == 0 {
		return 0
	}
	// count
	cnt := map[int]int{}
	for _, v := range power {
		cnt[v]++
	}
	// unique sorted values and gains
	vals := make([]int, 0, len(cnt))
	for v := range cnt {
		vals = append(vals, v)
	}
	sort.Ints(vals)
	gain := make([]int64, len(vals))
	for i, v := range vals {
		gain[i] = int64(v) * int64(cnt[v])
	}

	// dp with two pointers:
	// dp[i] = max(dp[i-1], dp[j] + gain[i-1]), где vals[i-1] - vals[j-1] >= 3
	dp := make([]int64, len(vals)+1)
	j := 0
	for i := 1; i <= len(vals); i++ {
		// сдвигаем j максимально вправо, сохраняя vals[i-1]-vals[j] >= 3
		for j < i-1 && vals[i-1]-vals[j] >= 3 {
			j++
		}
		// j указывает на первый индекс, который НЕ удовлетворяет условию,
		// значит совместимый последний — j-1, а его dp это dp[j]
		// (т.к. dp — на 1 больше по индексации)
		choose := dp[j] + gain[i-1]
		if dp[i-1] > choose {
			dp[i] = dp[i-1]
		} else {
			dp[i] = choose
		}
	}
	return dp[len(vals)]
}
