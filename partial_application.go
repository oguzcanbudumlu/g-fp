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

	// currying
	fmt.Println(threeSum(10, 20, 30))
	fmt.Println(threeSumCurried(10)(20)(30))
	fmt.Println(DogSpawnerCurry(Havanese)(Male)("Bucksy"))

	// server
	server := NewServer(MaxConnection(10), ServerName("MyFirstServer"))
	fmt.Printf("%+v\n", server)
}

// spawner

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

	havaneseSpawner          = DogSpawnerCurry(Havanese)
	poodleSpawner            = DogSpawnerCurry(Poodle)
	maleHavaneseSpawnerCurry = havaneseSpawner(Male)
	femalePoodleSpawnerCurry = poodleSpawner(Female)
)

// currying
func threeSum(a, b, c int) int {
	return a + b + c
}

func threeSumCurried(a int) func(int) func(int) int {
	return func(b int) func(int) int {
		return func(c int) int {
			return a + b + c
		}
	}
}

func DogSpawnerCurry(breed Breed) func(Gender) func(Name) Dog {
	return func(gender Gender) func(Name) Dog {
		return func(name Name) Dog {
			return Dog{
				Breed:  breed,
				Gender: gender,
				Name:   name,
			}
		}
	}
}

// Server Constructor

type (
	ServerOptions func(options) options
	TransportType int
)

const (
	UDP TransportType = iota
	TCP
)

type Server struct {
	options
	isAlive bool
}

type options struct {
	MaxConnection int
	TransportType TransportType
	Name          string
}

func MaxConnection(n int) ServerOptions {
	return func(o options) options {
		o.MaxConnection = n
		return o
	}
}

func ServerName(n string) ServerOptions {
	return func(o options) options {
		o.Name = n
		return o
	}
}

func Transport(t TransportType) ServerOptions {
	return func(o options) options {
		o.TransportType = t
		return o
	}
}

func NewServer(os ...ServerOptions) Server {
	opts := options{
		TransportType: TCP,
	}
	for _, option := range os {
		opts = option(opts)
	}

	return Server{
		options: opts,
		isAlive: true,
	}
}
