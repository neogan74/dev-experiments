package main

import (
	"fmt"
)

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	var m int64
	_, err := fmt.Scan(&m)
	if err != nil {
		panic(err)
	}

	const INF int64 = 1e18

	a := make([]int64, 31)
	for i := 0; i <= 30; i++ {
		_, err = fmt.Scan(&a[i])
		if err != nil {
			panic(err)
		}
	}

	// Улучшаем цены: a[i] = min(a[i], a[i-1]*2)
	for i := 1; i <= 30; i++ {
		a[i] = min(a[i], a[i-1]*2)
	}

	answer := INF
	var sum int64 = 0

	for i := 30; i >= 0; i-- {
		pow := int64(1) << i
		cnt := m / pow
		sum += cnt * a[i]
		m = m % pow
		// Рассматриваем возможность докупить 1 карточку уровня i и покрыть весь остаток
		if m > 0 {
			answer = min(answer, sum+a[i])
		} else {
			answer = min(answer, sum)
		}
	}

	fmt.Println(answer)
}
