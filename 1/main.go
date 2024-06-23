package main

import "fmt"

func main() {
	defer handlePanic()

	friends := []string{
		"Abylay",
		"Aigali",
		"Nur-Assyl",
	}
	fmt.Println("before panic")
	fmt.Println(friends)
	friends[3] = "Tamerlan"
	fmt.Println(friends)       // not showed
	fmt.Println("after panic") //not showed

}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
	fmt.Println("Panic handled succesfully!")
}
