package numberofwaystodividealongcorridor

// numberOfWays returns the number of ways to split the corridor into sections
// with exactly two seats each. The formula multiplies the gaps between seat
// pairs; see Readme.md for details.
func numberOfWays(corridor string) int {
	const mod = 1_000_000_007

	var seats []int
	for i, c := range corridor {
		if c == 'S' {
			seats = append(seats, i)
		}
	}

	if len(seats) == 0 || len(seats)%2 == 1 {
		return 0
	}

	ans := 1
	for i := 2; i < len(seats); i += 2 {
		gap := seats[i] - seats[i-1]
		ans = (ans * gap) % mod
	}

	return ans
}
