package _135_candy

func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)

	// Изначально каждому ребенку хотя бы 1 конфету
	for i := range candies {
		candies[i] = 1
	}

	// Первый проход слева направо
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// Второй проход справа налево
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}

	// Подсчитаем сумму всех конфет
	total := 0
	for _, c := range candies {
		total += c
	}

	return total
}
