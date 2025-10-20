package _011_FinalValueofVariableAfterPerformingOperations

func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, op := range operations {
		if op[1] == '+' { // "++X" или "X++"
			x++
		} else { // "--X" или "X--"
			x--
		}
	}
	return x
}
