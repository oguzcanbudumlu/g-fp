package main

import "fmt"

type predicate3 func(int) bool

func filter5(is []int, condition predicate3) []int {
	var out []int
	for _, i := range is {
		if condition(i) {
			out = append(out, i)
		}
	}

	return out
}

func createLargerThanPredicate(threshold int) predicate3 {
	return func(i int) bool {
		return i > threshold
	}
}

func main() {
	ints := []int{1, 2, 3}
	largerThanTwo2 := createLargerThanPredicate(2)

	fmt.Printf("%v", filter5(ints, largerThanTwo2))
}
