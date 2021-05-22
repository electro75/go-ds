package main

import (
	"fmt"
	"net/http"
	"os"
)

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

func httpreq() {
	res, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999) // the number indicates the number of spaces available in the slice initially.
	res.Body.Read(bs)
	fmt.Println(string(bs))
}
