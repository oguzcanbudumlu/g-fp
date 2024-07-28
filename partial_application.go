package main

import "fmt"

func createGreeting2(greeting string) func(string) string {
	return func(name string) string {
		return greeting + name
	}
}

func main() {
	firstGreeting := createGreeting2("Hello there")
	secondGreeting := createGreeting2("Hola")
	fmt.Println(firstGreeting("Remi"))
	fmt.Println(secondGreeting("Seen"))

	// spawner
	bucky := maleHavaneseSpawner("Bucky")
	rocky := maleHavaneseSpawner("Rocky")
	tipsy := femalePoodleSpawner("tipsy")
	fmt.Printf("%v\n%v\n%v\n", bucky, rocky, tipsy)
}

type (
	Name        string
	Breed       int
	Gender      int
	NameToDogFn func(Name) Dog
)

// Possible Breeds
const (
	Bulldog Breed = iota
	Havanese
	Cavalier
	Poodle
)
const (
	Male Gender = iota
	Female
)

type Dog struct {
	Name   Name
	Breed  Breed
	Gender Gender
}

func createDogsWithoutParitalApplication() []Dog {
	return []Dog{
		{Name: "Bucky", Breed: Havanese, Gender: Male},
		{Name: "Rocky", Breed: Havanese, Gender: Male},
		{Name: "Tipsy", Breed: Poodle, Gender: Female},
	}
}

func DogSpawner(breed Breed, gender Gender) NameToDogFn {
	return func(name Name) Dog {
		return Dog{Breed: breed, Gender: gender, Name: name}
	}
}

var (
	maleHavaneseSpawner = DogSpawner(Havanese, Male)
	femalePoodleSpawner = DogSpawner(Poodle, Female)
)
