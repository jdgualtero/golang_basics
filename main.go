package main

import "fmt"
import "./variables"
import "./arrays"
import "./cases"
import "./forpakage"

func main()  {

	fmt.Println("Call examples variables...")
	fmt.Println()
	variables.ShowVariables()
	arrays.ExampleArrays()
	cases.SwitchExample()
	forpakage.ForExample()
}