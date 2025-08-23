package largest3samedigitnumberinstring

func largestGoodInteger(num string) string {
	max := ""
	for i := 0; i <= len(num)-3; i++ {
		triple := num[i : i+3]
		if triple[0] == triple[1] && triple[1] == triple[2] {
			if triple > max {
				max = triple
			}
		}
	}
	return max
}
