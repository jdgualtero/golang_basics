package arrays

import "fmt"

func ExampleArrays() {
	// declare some arrays

	var stringArray [5]string
	stringArray[0] = "Hello"
	stringArray[1] = "World"
	stringArray[2] = "Example"

	fmt.Println("Array Example", stringArray)
	
	// iniatialize an array literal in go

	numberArray := [5]int{1,2,3,4,5}
	
	// iterate over the array
	for index :=0; index < len(numberArray); index++ {
		fmt.Println("number",numberArray[index])
	}

	// copy array

	copyIntArray := numberArray

	fmt.Println("numberArray",numberArray)
	fmt.Println("copyIntArray",copyIntArray)


	// trick with arrays

	fmt.Println("copyIntArray[0:2]",copyIntArray[0:2])
	fmt.Println("copyIntArray[2:]",copyIntArray[2:])

}