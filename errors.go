package main

import (
	"errors"
	"fmt"
)

var (
	validationErr = errors.New("validation err")
	notfoundErr   = errors.New("not found err")
)

func GetByID(id string) error {
	if id == "" {
		return validationErr
	}

	if id != "ivan" {
		return notfoundErr
	}

	return nil
}

func main() {

	err := GetByID("ivan")
	if err != nil {
		if errors.Is(err, validationErr) {
			fmt.Println("validation err")
		} else if errors.Is(err, notfoundErr) {
			fmt.Println("notfound err")
		} else {
			fmt.Println("unknown err")
		}
	}

}
