package interaction

func Repeat(char string) (repeated string) {
	for i := 0; i < 5; i++ {
		repeated = repeated + char
	}
	return
}
