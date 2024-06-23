package main

import "fmt"

func main() {
	defer printMeassage() //3
	defer greeting()      // 2
	fmt.Println("main()") // 1
	// friends := []string{
	// 	"Abylay",
	// 	"Aigali",
	// 	"Nur-Assyl",
	// }
	// friends[3] = "Tamerlan"

	// fmt.Println(friends)
	panic("i get a fuck") // last or 4
}

func printMeassage() {
	fmt.Println("printMessage()")
}

func greeting() {
	fmt.Println("Hello from greeting()")
}
