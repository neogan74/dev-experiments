package _838_push_domino

func pushDominoes(dominoes string) string {
	type symbol struct {
		index int
		char  byte
	}
	// Identify positions of all non-dot dominoes ('L' or 'R')
	symbols := []symbol{}
	for i := 0; i < len(dominoes); i++ {
		if dominoes[i] != '.' {
			symbols = append(symbols, symbol{i, dominoes[i]})
		}
	}

	// Add virtual boundary dominoes to simplify edge cases
	symbols = append([]symbol{{-1, 'L'}}, symbols...)
	symbols = append(symbols, symbol{len(dominoes), 'R'})

	// Convert the string
	ans := []byte(dominoes)

	// Process each pair of adjacent symbols
	for i := 0; i < len(symbols)-1; i++ {
		left := symbols[i]
		right := symbols[i+1]
		// Case same direction
		if left.char == right.char {
			for k := left.index + 1; k < right.index; k++ {
				ans[k] = left.char
			}
		} else if left.char == 'R' && right.char == 'L' {
			l, r := left.index+1, right.index-1
			for l < r {
				ans[l] = 'R'
				ans[r] = 'L'
				l++
				r--
			}
		}
	}

	return string(ans)
}
