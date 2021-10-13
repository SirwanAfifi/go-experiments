package main

import "fmt"

// https://www.calhoun.io/one-liner-if-statements-with-errors/
func main() {
	var name string = "Sirwan";

	if username := name + "123"; username == "Sirwan123" {
		fmt.Println(username)
	}
}