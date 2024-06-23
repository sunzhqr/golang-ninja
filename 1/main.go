package main

import "fmt"

func main() {
	users := map[string]int{
		"Sanzhar":   19,
		"Abylay":    19,
		"Nur-Assyl": 20,
	}

	AigaliExists(users)
	users["Aigali"] = 19
	AigaliExists(users)

	fmt.Println("Before: ", len(users))

	for k, v := range users {
		fmt.Printf("%s, %d deleted!\n", k, v)
		delete(users, k)
	}

	fmt.Println("After: ", len(users))

}

func AigaliExists(users map[string]int) {
	age, exists := users["Aigali"]
	if exists {
		fmt.Println("Aigali", age)
	} else {
		fmt.Println("Where is Aigali??!")
	}
}
