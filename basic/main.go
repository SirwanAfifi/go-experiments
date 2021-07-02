package main

import (
	"fmt"
	"strings"
)

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

	var newSkills = append(userA.Skills, "Java")

	fmt.Println(newSkills[len(newSkills)-1])
	fmt.Println(newSkills[2:]) // Range newSkills[1:3]

	greeting := "Hello there"
	fmt.Println(strings.Contains(greeting, "Hello"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "hi"))

	users := map[int]Person{
		1: {
			Id:     1,
			Name:   "Sirwan",
			Age:    32,
			Skills: []string{"JS", "Go", "C#", "Kotlin", "Swift"},
		},
	}

	for k, v := range users {
		fmt.Println(k, v)
	}

	phonebooks := map[int]string{
		1232434: "Sirwan",
		3434454: "UserA",
	}

	for k, v := range phonebooks {
		fmt.Println(k, v)
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
