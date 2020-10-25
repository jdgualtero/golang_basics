package main

import "fmt"
import "./variables"
import "./arrays"
import "./cases"

func main()  {

	fmt.Println("Call examples variables...")
	fmt.Println()
	variables.ShowVariables()
	arrays.ExampleArrays()
	cases.SwitchExample()
}