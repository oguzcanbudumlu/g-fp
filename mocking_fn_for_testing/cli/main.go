package main

import "g-fp/mocking_fn_for_testing/pkg"

func main() {
	t := pkg.NewTodo()
	t.Write("hi")
}
