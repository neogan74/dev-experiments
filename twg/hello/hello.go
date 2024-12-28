package hello

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "GO"
	}
	if language == "Spanish" {
		return spanishHelloPrefix + name + "!"
	}
	if language == "French" {
		return frenchHelloPrefix + name + "!"
	}
	return englishHelloPrefix + name + "!"
}
