package main

import (
	"fmt"
	"slices"
)

func main() {
	name := []string{"ivan", "satria", "muhammad", "joko", "khanedy"}
	values := []int{100, 90, 80, 75}

	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(name, "satria"))
	fmt.Println(slices.Index(name, "muhammad"))
}
