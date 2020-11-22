package map_package

import ("fmt")

func BasicExamples() {
	// firs example

	fmt.Println("Firts Map")

	var firts_map = map[string]int {"Jose":10, "Maria": 11}
	
	fmt.Println(firts_map)


	fmt.Println("Empty map")

	var empty_map = map[string]int{}
	fmt.Println(empty_map)

	fmt.Println("Empty map created with make function")

	var map_by_make = make(map[string]int)
	map_by_make["sofia"] = 9
	map_by_make["miguel"] = 8

	fmt.Println(map_by_make)

	fmt.Println("Accesing Items")
	fmt.Println(map_by_make["sofia"]) 

	// len of a map

	fmt.Println("Length of map created with Make.",len(map_by_make))

	fmt.Println("Add items in a map")

	map_by_make["anabel"] = 7

	fmt.Println("Length of map created with Make.",len(map_by_make))

	// update items in a map

	fmt.Println(map_by_make) 
	map_by_make["anabel"] = 9
	fmt.Println(map_by_make)

	// delete items in a map

	delete(map_by_make, "anabel")
	fmt.Println(map_by_make)

}