package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("hello/wordl.go"))
	fmt.Println(path.Base("hello/wordl.go"))
	fmt.Println(path.Ext("hello/wordl.go"))
	fmt.Println(path.Join("hello", "wordl", "main.go"))
}
