package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

// logic without interfaces
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sp spanishBot) {
// 	fmt.Println((sp.getGreeting()))
// }

// with interface
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// if the received value is not being used in the function,
// it can be ommitted completely
func (englishBot) getGreeting() string {
	return "Hello, apples"
}

func (sp spanishBot) getGreeting() string {
	return "Hola, manzanas"
}
