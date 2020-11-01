package forpakage

import (
	"fmt"
)

func ForExample() {
	fmt.Println("Basic for")

	index := 0

	for ;index <= 10; index++ {
		fmt.Println("index", index)
	}


	// Example for range Statement
	strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
	for index, element := range strDict {
		fmt.Println("Index :", index, " Element :", element)
	}
 
	// Example 2
	for key := range strDict {
		fmt.Println(key)
	}
 
	// Example 3
	for _, value := range strDict {
		fmt.Println(value)
	}

	// infinite loop

	limit := 0;

	for {
		if limit > 10 {
			break
		} else {
			fmt.Println("limit", limit)
			limit++
		}
	}
}