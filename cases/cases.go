package cases

import (
	"fmt"
	"time"
)

func SwitchExample() {

	today := time.Now().Weekday()

	switch int(today) {
	case 0:
		fmt.Println("Today is sunday")
	case 1:
		fmt.Println("Today is monday")
	case 2:
		fmt.Println("Today is Tuesday")
	case 3:
		fmt.Println("Today is Wednesday")
	case 4:
		fmt.Println("Today is Thursday")
	case 5:
		fmt.Println("Today is Friday")
	case 6:
		fmt.Println("Today is Saturday")
	
	default:
		fmt.Println("is other day",today)
		
	}
	
	// switch multiple values

	today2 := time.Now().Day()

	switch today2 {
	case 15,30:
		fmt.Println("you are goint to receive money.")
	case 10,20,24:
		fmt.Println("you must to pay internet subscription.")
	
	default:
		fmt.Println("is other day")
		
	}

	today3 := time.Now()
 
	switch today3.Day() {
	case 5:
		fmt.Println("Clean your house.")
		fallthrough
	case 10:
		fmt.Println("Buy some wine.")
		fallthrough
	case 15:
		fmt.Println("Visit a doctor.")
		fallthrough
	case 24:
		fmt.Println("Buy some food.")
		fallthrough
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}

}