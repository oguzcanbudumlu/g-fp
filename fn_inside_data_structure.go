package main

import "fmt"

type predicate4 func(int) bool

func createLargerThanPredicate4(threshold int) predicate4 {
	return func(i int) bool {
		return i > threshold
	}
}
func filter_(is []int, condition predicate4) []int {
	var out []int
	for _, i := range is {
		if condition(i) {
			out = append(out, i)
		}
	}

	return out
}

var (
	largerThanTwo_     = createLargerThanPredicate4(2)
	largerThanFive_    = createLargerThanPredicate4(5)
	largerThanHundred_ = createLargerThanPredicate4(100)
)

func main() {
	ints := []int{1, 2, 3, 6, 10}
	predicates := []predicate4{largerThanTwo_, largerThanFive_, largerThanHundred_}

	for _, predicate_ := range predicates {
		fmt.Printf("%v\n", filter_(ints, predicate_))
	}

	// fn inside dt
	dispatcher := map[string]predicate4{
		"2": largerThanTwo_,
		"5": largerThanFive_,
	}

	fmt.Printf("%v\n", filter_(ints, dispatcher["2"]))
}
