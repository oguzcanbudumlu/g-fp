package main

import "fmt"

type phoneNumber string
type age int

func (a age) valid() bool {
	return a < 120
}

func isValidPerson(p Person2) bool {
	return p.age.valid() && p.name != ""
}

type Person2 struct {
	name        string
	phonenumber phoneNumber
	age         age
}

func (p *Person2) setPhoneNumber(s phoneNumber) {
	p.phonenumber = s
}

func main() {
	p := Person2{
		name:        "John",
		phonenumber: "123",
	}

	fmt.Printf("%v\n", p)
}
