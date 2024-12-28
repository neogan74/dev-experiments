package hello

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "GO"
	}
	if language == "Spanish" {
		return spanishHelloPrefix + name + "!"
	}
	return englishHelloPrefix + name + "!"
}
