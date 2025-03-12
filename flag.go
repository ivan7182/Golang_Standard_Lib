package main

import (
	"flag"
	"fmt"
)

func main() {
	var username = flag.String("username", "root", "database username")
	var password = flag.String("password", "root", "database password")
	var host = flag.String("host", "root", "database host")
	var port = flag.Int("post", 0, "database port")

	flag.Parse()

	fmt.Println("username", *username)
	fmt.Println("password", *password)
	fmt.Println("host", *host)
	fmt.Println("port", *port)
}
