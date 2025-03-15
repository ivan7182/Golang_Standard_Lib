package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("hello wordl\n")
	_, _ = writer.WriteString("selamat belajar\n")

	writer.Flush()
}
