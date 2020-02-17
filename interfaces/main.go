package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eng := englishBot{}
	span := spanishBot{}

	printGreeting(eng)
	printGreeting(span)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
