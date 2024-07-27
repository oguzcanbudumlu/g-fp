package main

import (
	"errors"
	"fmt"
)

func main() {
	greetingFn := createGreeting()
	response := greetingFn("Ana")
	fmt.Println(response)
}

func fn() (string, error) {
	return "", errors.New("error 1")
}

func fn2() (string, error) {
	return "", errors.New("error 2")
}

func outerFn() func() {
	fmt.Println("outer fn")
	return func() {
		fmt.Println("inner fn")
	}
}

func createGreeting() func(string) string {
	s := "Hello "
	return func(name string) string {
		return s + name
	}
}
