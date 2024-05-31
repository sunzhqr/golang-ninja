package main

import "fmt"

func main() {
	message, _ := enterTheClub(7)
	fmt.Println(message)
	//fmt.Println(entered)
	//fmt.Println(enterTheClub(17))
}

func enterTheClub(age int) (string, bool) {
	if age >= 18 && age < 45 {
		return "You accepted a club!", true
	} else if age >= 45 && age <= 65 {
		return "Are you sure?", true
	} else if age >= 65 {
		return "Not for you!", false
	}
	return "You cannot enter to the Club(", false
}
