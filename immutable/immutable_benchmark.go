package immutable

type Human struct {
	name string
	age  int
}

func ImmutableCreateHuman() Human {
	h := Human{}
	h = immutableSetName(h, "Sean")
	h = immutableSetAge(h, 20)
	return h
}

func immutableSetName(h Human, name string) Human {
	h.name = name
	return h
}

func immutableSetAge(h Human, age int) Human {
	h.age = age
	return h
}

func MutableCreateHuman() *Human {
	h := &Human{}
	mutableSetName(h, "Tom")
	mutableSetAge(h, 30)
	return h
}

func mutableSetName(h *Human, name string) {
	h.name = name
}

func mutableSetAge(h *Human, age int) {
	h.age = age
}
