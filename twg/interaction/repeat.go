package interaction

func Repeat(char string, rep int) (repeated string) {
	for i := 0; i < rep; i++ {
		repeated += char
	}
	return
}
