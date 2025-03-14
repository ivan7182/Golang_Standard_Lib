package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
}

type Sample struct {
	Name string `required :"true" max:"10"`
}

func ReadField(value any) {
	ValueType := reflect.TypeOf(value)
	fmt.Println("type name", ValueType.Name())
	for i := 0; i < ValueType.NumField(); i++ {
		structField := ValueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))

	}
}

func IsValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}

	}

	return result
}

func main() {
	ReadField(Person{"ivan", "", "muh@gmail.com"})
	ReadField(Sample{"satria"})

	person := Person{
		Name:    "ivan",
		Address: "",
		Email:   "mu@gmail.com",
	}

	fmt.Println(IsValid(person))
}
