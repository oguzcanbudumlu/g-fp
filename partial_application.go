package main

import "fmt"

func createGreeting2(greeting string) func(string) string {
	return func(name string) string {
		return greeting + name
	}
}

func main() {
	firstGreeting := createGreeting2("Hello there")
	secondGreeting := createGreeting2("Hola")
	fmt.Println(firstGreeting("Remi"))
	fmt.Printf(secondGreeting("Seen"))
}
