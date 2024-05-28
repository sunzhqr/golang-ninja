package main

import "fmt"

func main() { //Entry point in application
	a, b, c := 1, 2, true
	fmt.Println(a, b, c)
	a, b = b, a
	fmt.Println(a, b, c)
	a, c = 10, false
	fmt.Println(a, b, c)
}
