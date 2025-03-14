package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	value := "mtk, BI, kimia\n" +
		"penjas, fisika, agama\n" +
		"olahraga, ppkn, indonesia"

	reader := csv.NewReader(strings.NewReader(value))
	for {
		recorde, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(recorde)
	}
}
