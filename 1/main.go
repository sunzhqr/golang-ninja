package main

import "fmt"

func main() {
	message, _ := enterTheClub(17)
	fmt.Println(message)
	//fmt.Println(entered)
	//fmt.Println(enterTheClub(17))
}

func enterTheClub(age int) (string, bool) {
	if age >= 18 {
		return "You accepted a club!", true
	} else {
		return "You cannot enter to the Club(", false
	}
}
