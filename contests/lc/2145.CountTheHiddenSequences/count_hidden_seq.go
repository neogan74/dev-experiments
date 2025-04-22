package _145_CountTheHiddenSequences

// countHiddenSequences возвращает количество скрытых последовательностей,
// длина которых = len(differences)+1, все элементы ∈ [lower, upper],
// и для которых differences[i] = hidden[i+1] - hidden[i].
func countHiddenSequences(differences []int, lower, upper int) int {
	prefix := 0
	minPref, maxPref := 0, 0

	// вычисляем префиксную сумму и одновременно min/max
	for _, d := range differences {
		prefix += d
		if prefix < minPref {
			minPref = prefix
		}
		if prefix > maxPref {
			maxPref = prefix
		}
	}

	// допустимый сдвиг x должен лежать в
	// [lower - minPref, upper - maxPref]
	// число целых в этом отрезке:
	total := (upper - maxPref) - (lower - minPref) + 1
	if total < 0 {
		return 0
	}
	return total
}
