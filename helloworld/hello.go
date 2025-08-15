package main

import "fmt"

func Hello(str string, language string) string {

	prefix := GetPrefix(language)

	if str == "" {
		return prefix + "World"
	}
	return prefix + str
}
func GetPrefix(lang string) (prefix string) {

	engPrefix := "Hello "
	frenchPrefix := "Bonjour "
	spanishPrefix := "Hola "

	if lang == "french" {
		prefix = frenchPrefix
	}
	if lang == "spanish" {
		prefix = spanishPrefix
	}
	if lang == "english" {
		prefix = engPrefix
	}
	return

}

func main() {
	fmt.Println(Hello("Hello ", "gg"))
}
