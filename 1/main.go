package main

import "fmt"

type User struct {
	name   string
	age    int
	gender string
}

// Value Receiver or Getter
func (u User) getName() string {
	fmt.Println(u.name)
	return u.name
}

// Pointer Receiver or Setter
func (u *User) setName(name string) {
	u.name = name
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
	user1.getName()
	user2.getName()
	user1.setName("Example1")
	user2.setName("Example2")
	user1.getName()
	user2.getName()
	fmt.Println(user1)
}
