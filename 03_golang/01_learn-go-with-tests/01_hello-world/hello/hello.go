package hello

import (
	"fmt"
)

const (
	englishLanguage    = "english"
	spanishLanguage    = "spanish"
	frenchLanguage     = "french"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func main() {
	name := "World"
	fmt.Println(Hello(name, "english"))
}

// Hello prints greeting
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s%s", greeting(language), name)
}

func greeting(language string) string {
	greeting := ""

	switch language {
	case spanishLanguage:
		greeting = spanishHelloPrefix

	case frenchLanguage:
		greeting = frenchHelloPrefix

	default:
		greeting = englishHelloPrefix
	}

	return greeting
}
