package main

import "fmt"

type predicate5 func(int) bool

type ConstraintChecker struct {
	largerThan  predicate5
	smallerThan predicate5
}

func createLargerThanPredicate5(threshold int) predicate5 {
	return func(i int) bool {
		return i > threshold
	}
}

func (c ConstraintChecker) check(input int) bool {
	return c.largerThan(input) && c.smallerThan(input)
}

func main() {
	checker := ConstraintChecker{
		largerThan:  createLargerThanPredicate5(2),
		smallerThan: func(i int) bool { return i < 10 },
	}
	fmt.Printf("%v\n", checker.check(5))
}
