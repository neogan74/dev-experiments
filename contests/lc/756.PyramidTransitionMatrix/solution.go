package main

func pyramidTransition(bottom string, allowed []string) bool {
	pairs := make(map[string][]byte)
	for _, rule := range allowed {
		pair := rule[:2]
		pairs[pair] = append(pairs[pair], rule[2])
	}

	memo := make(map[string]bool)

	var canBuild func(row string) bool
	canBuild = func(row string) bool {
		if len(row) == 1 {
			return true
		}
		if val, ok := memo[row]; ok {
			return val
		}

		next := make([]byte, len(row)-1)
		var buildNext func(pos int) bool
		buildNext = func(pos int) bool {
			if pos == len(row)-1 {
				return canBuild(string(next))
			}

			pair := row[pos : pos+2]
			choices, ok := pairs[pair]
			if !ok {
				return false
			}
			for _, ch := range choices {
				next[pos] = ch
				if buildNext(pos + 1) {
					return true
				}
			}
			return false
		}

		res := buildNext(0)
		memo[row] = res
		return res
	}

	return canBuild(bottom)
}
