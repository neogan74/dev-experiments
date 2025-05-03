package _1007_min_domino_rotation_for_eq_row

func minDominoRotations(tops []int, bottoms []int) int {
	cnta := make([]int, 7)
	cntb := make([]int, 7)
	same := make([]int, 7)
	for i := range tops {
		cnta[tops[i]]++
		cntb[bottoms[i]]++
		if tops[i] == bottoms[i] {
			same[tops[i]]++
		}
	}
	for i := 1; i <= 6; i++ {
		if cnta[i]+cntb[i]-same[i] == len(tops) {
			if cnta[i] < cntb[i] {
				return cnta[i] - same[i]
			}
			return cntb[i] - same[i]
		}
	}
	return -1

}
