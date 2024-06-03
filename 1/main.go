package main

import (
	"errors"
	"fmt"
)

func main() {
	response, err := predictionForAbylay("Saturday")
	if err != nil {
		fmt.Println("Error", "WTF, you can't enter the day of Week?")
	} else {
		fmt.Println(response)
	}
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

func predictionForAbylay(dayOfWeek string) (string, error) {
	switch dayOfWeek {
	case "Monday":
		return "Statistics fx", nil
	case "Tuesday":
		return "", nil
	case "Wednesday":
		return "Academic fx", nil
	case "Thursday":
		return "AP fx", nil
	case "Friday":
		return "Alem exam lose", nil
	case "Saturday":
		return "Military double dive into Abylay", nil
	case "Sunday":
		return "Othodnyak", nil
	default:
		return "Abylay wins", errors.New("Err")
	}
}
