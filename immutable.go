package main

import "fmt"

type Human struct {
	name string
	age  int
}

func setName(h Human, name string) Human {
	h.name = name
	return h
}

func setNameP(h *Human, name string) {
	h.name = name
}

func main() {
	p := Human{
		name: "Benny",
		age:  55,
	}

	setName(p, "Bjorn")
	fmt.Println(p.name)

	setNameP(&p, "Bjorn")
	fmt.Println(p.name)

	p = setName(p, "Kamil")
	fmt.Println(p.name)

	// collection
	names := []string{"Miranda", "Paula"}
	names = append(names, "Ahmet")
	fmt.Printf("%v\n", names)

	// map
	m := map[string]int{}
	addValue(m, "red", 10) // always pass-by-ref
	fmt.Printf("%v\n", m)

	// slice
	ns := []string{"Miranda"}
	addValueS(ns, "Mehmet")
	fmt.Printf("%v\n", ns)

	// slice pointer
	nsp := []string{"Miranda"}
	addValueSP(&nsp, "Mehmet")
	fmt.Printf("%v\n", nsp)
}

func addValue(m map[string]int, colour string, value int) {
	m[colour] = value
}

func addValueS(s []string, name string) {
	s = append(s, name)
}

func addValueSP(s *[]string, name string) {
	*s = append(*s, name)

}
