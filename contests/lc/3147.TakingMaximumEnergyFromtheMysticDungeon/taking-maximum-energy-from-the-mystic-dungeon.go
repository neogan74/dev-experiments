package taking_maximum_energy_from_the_mystic_dungeon

func maxEnergy(energy []int, k int) int {
	n := len(energy)
	best := -1 << 60
	limit := k
	if n < k {
		limit = n
	}
	for r := 0; r < limit; r++ {
		// последний индекс с остатком r
		last := ((n-1-r)/k)*k + r
		sum := 0
		for j := last; j >= 0; j -= k {
			sum += energy[j]
			if sum > best {
				best = sum
			}
		}
	}
	return best
}
