package main

import (
	"fmt"
	"reflect"
)

func main() {
	type anotherStructure struct {
		Field string
	}

	type sampleStructure struct {
		Name   string
		Age    int
		Height float32

		Another anotherStructure
	}

	sType := reflect.TypeOf(sampleStructure{})

	for i := 0; i < sType.NumField(); i++ {
		field := sType.Field(i)

		fmt.Printf("%s -> %s\n", field.Name, field.Type.Name())
	}
}
