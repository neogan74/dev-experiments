package hello

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "GO"
	}
	return englishHelloPrefix + name + "!"
}
