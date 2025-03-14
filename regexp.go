package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`i([a-z])([a-z])n`)

	fmt.Println(regex.MatchString("ivan"))
	fmt.Println(regex.MatchString("satria"))
	fmt.Println(regex.MatchString("elsa"))

	fmt.Println(regex.FindAllString("ivan ipan Ipan", 10))

}
