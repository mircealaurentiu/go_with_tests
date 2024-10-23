package main

import "fmt"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := "Hello, "

	switch language {
	case "Spanish":
		prefix = "Hola, "
	case "Romanian":
		prefix = "Salut, "
	default:
		prefix = "Hello, "
	}

	return (prefix + name)
}

func main() {
	fmt.Println(Hello("Ioan", "Romanian"))
	fmt.Println(Hello("", ""))
}
