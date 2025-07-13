package minimumgeneticmutation

func minMutation(start string, end string, bank []string) int {
	bankSet := make(map[string]bool)
	for _, gene := range bank {
		bankSet[gene] = true
	}
	if !bankSet[end] {
		return -1 // Если конечного гена нет в банке — смысла нет
	}

	queue := []string{start}
	visited := map[string]bool{start: true}
	steps := 0
	genes := []byte{'A', 'C', 'G', 'T'}

	for len(queue) > 0 {
		for size := len(queue); size > 0; size-- {
			curr := queue[0]
			queue = queue[1:]

			if curr == end {
				return steps
			}

			// Пробуем менять по одной букве
			currBytes := []byte(curr)
			for i := 0; i < len(currBytes); i++ {
				original := currBytes[i]
				for _, g := range genes {
					currBytes[i] = g
					next := string(currBytes)
					if bankSet[next] && !visited[next] {
						visited[next] = true
						queue = append(queue, next)
					}
				}
				currBytes[i] = original // откат
			}
		}
		steps++
	}

	return -1
}
