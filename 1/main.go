package main

import "fmt"

type User struct {
	name   string
	age    int
	gender string
}

func NewUser(name, gender string, age int) User {
	return User{
		name:   name,
		age:    age,
		gender: gender,
	}
}

func main() {
	user1 := NewUser("Sanzhar", "Male", 19)
	user2 := User{"Aigali", 19, "Male"}
	fmt.Printf("%+v\n", user1.name)
	fmt.Printf("%+v\n", user2.age)
}
