package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	value := strings.NewReader("string panjang\ndengan line baru")
	read := bufio.NewReader(value)

	for {
		line, _, err := read.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}
}
