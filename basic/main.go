package main

import "fmt"

func main() {
	// name, age := getName()

	// fmt.Println(name, age)

	userA := Person{
		Id:     1,
		Name:   "Sirwan",
		Age:    32,
		Skills: []string{"JS", "Go", "C#", "Kotlin", "Swift"},
	}

	fmt.Println(userA.Skills)

	for _, value := range userA.Skills {
		fmt.Println(value)
	}
}

func getName() (string, int) {
	return "Sirwan", 32
}

type Person struct {
	Id     int
	Name   string
	Age    int
	Skills []string
}
