package variables


import "fmt"

func ShowVariables ()  {

	var variableInt int = 12
	var variableString string = "Hello world"
	var variableFloat  float64 = 3.14
	var variableBoolean bool = true

	fmt.Println("variable int", variableInt)
	fmt.Println("variable string", variableString)
	fmt.Println("variable float", variableFloat)
	fmt.Println("variable Boolean",variableBoolean)
}