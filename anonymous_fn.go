package main

func filter4(is []int, predicate func(int) bool) []int {
	out := []int{}
	for _, i := range is {
		if predicate(i) {
			out = append(out, i)
		}
	}
	return out
}

func main() {
	filter4([]int{1, 2, 3}, func(i int) bool { return i > 2 })
}
