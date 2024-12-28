package hello

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	defaultName := "GO"
	if name == "" {
		name = defaultName
	}
	return englishHelloPrefix + name + "!"
}
