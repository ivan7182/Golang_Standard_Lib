package main

import (
	"fmt"
	"strconv"
)

func main() {
	hasil, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(hasil)
	}

	hasilInt, err := strconv.Atoi("90")
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(hasilInt)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	var stringInt string = strconv.Itoa(999)
	fmt.Println(stringInt)

	floatStr, err := strconv.ParseFloat("3.14", 64)
	if err != nil {
		fmt.Println("err", err.Error())
	} else {
		fmt.Println(floatStr)
	}
}
