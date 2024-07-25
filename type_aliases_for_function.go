package main

type predicate2 func(int) bool

func largerThanTwo(i int) bool {
	return i > 2
}
func filter2(is []int, predicate func(int) bool) []int {
	out := []int{}
	for _, i := range is {
		if predicate(i) {
			out = append(out, i)
		}
	}
	return out
}

func main() {
	ints := []int{1, 2, 3}
	filter2(ints, largerThanTwo)
}
