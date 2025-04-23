package _81_RabbitsinForest

import "fmt"

// numRabbits возвращает минимальное число кроликов в лесу,
// зная для каждого ответа answers[i] = сколько других кроликов того же цвета видел i‑й.
func numRabbits(answers []int) int {
	counts := make(map[int]int)
	for _, ans := range answers {
		counts[ans]++
	}

	total := 0
	for ans, freq := range counts {
		// если кролик отвечает ans, то всего таких в группе ans+1
		groupSize := ans + 1
		// сколько полных групп нужно, чтобы покрыть freq ответов
		groups := (freq + groupSize - 1) / groupSize
		total += groups * groupSize
	}

	return total
}

func main() {
	fmt.Println(numRabbits([]int{1, 1, 2}))          // 5
	fmt.Println(numRabbits([]int{0, 0, 1, 1, 1}))    // 6
	fmt.Println(numRabbits([]int{2, 2, 2, 2, 2, 2})) // 6  (две группы по 3)
}
