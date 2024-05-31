package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	message, err := enterTheClub(27)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

func enterTheClub(age int) (string, error) {
	if age >= 18 && age < 45 {
		return "You accepted a club!", nil
	} else if age >= 45 && age <= 65 {
		return "Are you sure?", nil
	} else if age >= 65 {
		return "Not for you!", errors.New("you are too old!")
	}
	return "You cannot enter to the Club(", errors.New("you are too young!")
}
