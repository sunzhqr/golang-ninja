package main

import (
	"fmt"
	"reflect"
)

func main() { //Entry point in application
	//Initialization and assignment value of variable
	message1 := "Hello, GO!" //variable := "something"

	message2 := ""
	// Assign value
	message2 = "Hello, Golang!"

	var message3 string
	message3 = "Hello, Golang!"

	// Initialization
	var num int
	fmt.Println(reflect.TypeOf(num))
	//Constant variable
	const Pi float64 = 3.1415926
	const name = "Sanzhar"
	fmt.Println(Pi)
	fmt.Println(message1, message2, message3)
}
