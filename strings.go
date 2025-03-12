package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("ivan satrian", "ivan"))
	fmt.Println(strings.Split("ivan satrian", " "))
	fmt.Println(strings.ToLower("ivan satrian"))
	fmt.Println(strings.ToUpper("ivan satrian"))
	fmt.Println(strings.Trim("      ivan ", " "))
	fmt.Println(strings.ReplaceAll("ivan satrian ivan januari", "ivan", "muhammad"))

}
