package main

import "fmt"

func main() {
	firstName := "ivan"
	lastName := "satria"

	fmt.Println("hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}
