package main

import "fmt"

const prefix = "Hello "

func Hello(str string) string {
	if str == "" {
		return "Hello World"
	}
	return prefix + str
}

func main() {
	fmt.Println(Hello("Hello "))
}
