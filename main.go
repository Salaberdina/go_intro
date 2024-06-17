package main

import "fmt"

const (
	englishHelloPrefix = "hello, "
	spanishHelloPrefix = "hola, "
	frenchHelloPrefix  = "bonjour, "
	havaiHelloPrefix   = "aloha, "

	spanish = "spanish"
	french  = "french"
	havai   = "havai"
)

func main() {
	fmt.Println(Hello("", ""))
}

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return getGreetingPrefix(language) + name
}

func getGreetingPrefix(language string) string {
	switch language {
	case french:
		return frenchHelloPrefix
	case spanish:
		return spanishHelloPrefix
	case havai:
		return havaiHelloPrefix
	default:
		return englishHelloPrefix
	}
}
