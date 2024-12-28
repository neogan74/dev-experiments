package hello

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "GO"
	}
	prefix := englishHelloPrefix
	switch language {
	case spanish:
		prefix = spanishHelloPrefix

	case french:
		prefix = frenchHelloPrefix
	}
	return prefix + name + "!"
}
