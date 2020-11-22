package main

import "fmt"
import "./variables"
import "./arrays"
import "./cases"
import "./forpakage"
import "./map_package"

func main()  {

	fmt.Println("Call examples variables...")
	fmt.Println()
	variables.ShowVariables()
	arrays.ExampleArrays()
	cases.SwitchExample()
	forpakage.ForExample()
	map_package.BasicExamples()

}