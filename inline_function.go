package main

import "fmt"

func filter3(is []int, predicate func(int) bool) []int {
	out := []int{}
	for _, i := range is {
		if predicate(i) {
			out = append(out, i)
		}
	}
	return out
}

func main() {
	inlinePersonStruct := struct {
		name string
	}{
		name: "John",
	}

	ints :=
		[]int{1, 2, 3}

	inlineFn := func(i int) bool { return i > 2 }

	filter3(ints, inlineFn)
	fmt.Printf("%v\n", inlinePersonStruct)
}
