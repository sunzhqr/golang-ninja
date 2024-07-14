package main

import "fmt"

type User struct {
	name   string
	age    int
	gender string
}

func main() {
	user1 := User{"Sanzhar", 19, "Male"}
	user2 := User{"Aigali", 19, "Male"}
	fmt.Printf("%+v\n", user1)
	fmt.Printf("%+v\n", user2)
}
