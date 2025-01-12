package main

import (
	"fmt"
	"reflect"
)

func main() {

	var name string = "jhon"
	fmt.Println("Name", name)

	var test = "test"
	fmt.Println("test", test)

	var a, c = 2, 3
	fmt.Println("a,c", a, c)

	var bol = true
	fmt.Println("boolean", bol)

	const country = "USA"
	fmt.Println("country", country)

	age := 30
	fmt.Println("age", age, " - Type of Age", reflect.TypeOf(age))

	NameStrin := "test string"
	fmt.Println("NameStrin ", NameStrin, "TYPE - ", reflect.TypeOf(NameStrin))

}
