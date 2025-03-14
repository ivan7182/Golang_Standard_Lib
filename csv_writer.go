package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"ivan", "satria", "muhammad"})
	_ = writer.Write([]string{"joko", "budi", "ibu"})
	_ = writer.Write([]string{"yanto", "resa", "arab"})

	writer.Flush()
}
